package engine

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/player"
)

func (e *Engine) tick(tic time.Time) {
	// TODO:metrics monitor ( tic - e.tic ) to see how regular our fanout procedure executes

	{
		e.tic = tic
	}

	// Fanout to every player participating in the game. Note that we are not
	// doing any unnecessary work if there are no players, because in this case
	// the memory table will be empty and the for loop will not iterate. This is
	// then also why the fanout cycle deadline does not have to be initialized in
	// engine.New().

	e.mem.ply.Range(func(k [2]byte, v *player.Player) bool {
		// Get the player specific buffer and reset it to start a new cycle.

		var b []byte
		{
			b, _ = e.buf.LoadAndDelete(k)
		}

		// Create a new timer for this fanout cycle. See Engine.join() for the
		// deadline calculation.

		var t *time.Timer
		{
			t = time.NewTimer(e.tim)
		}

		// Run the client specific fanout cycle based on a resource limited worker
		// pool that enforces the timely cadence of the globally required fanout
		// procedure.

		{
			go e.worker(v, b, t)
		}

		return true
	})
}

func (e *Engine) worker(ply *player.Player, buf []byte, tim *time.Timer) {
	// The semaphore below controls the amount of workers that are allowed to
	// process packets at the same time. Every time we process a player, we
	// allocate a ticket with the semaphore, before doing the actual work.

	{
		e.sem <- struct{}{}
	}

	// We use the early channel to stop processing before deadline violation.

	var ear chan struct{}
	{
		ear = make(chan struct{})
	}

	// Send the prepared fanout buffer to the given player. The use of the early
	// channel below allows us to proceed before our deadline expires. Note that
	// we don't check the error below, because Client.Stream() causes the client
	// connection to be terminated already in case there is an error produced.

	go func() {
		ply.Cli.Stream(buf) // nolint:errcheck
		close(ear)
	}()

	// Block this worker call until we are done processing the given player, or
	// until the provided deadline expired.

	select {
	case <-ear:
		tim.Stop()
	case <-tim.C:
	}

	// Ensure to release the semaphore that we allocated for this worker call.

	{
		<-e.sem
	}
}
