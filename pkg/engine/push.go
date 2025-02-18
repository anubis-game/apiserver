package engine

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/player"
)

func (e *Engine) push(tic time.Time) {
	// TODO:metrics monitor ( tic - e.tic ) to see how regular our fanout procedure executes

	{
		e.tic = tic
	}

	// Fanout to every player participating in the game. Note that we are not
	// doing any unnecessary work if there are no players, because in this case
	// the memory table will be empty and the for loop will not iterate. This is
	// then also why the fanout cycle deadline does not have to be initialized in
	// engine.New().

	for k, v := range e.mem.ply {
		// TODO:infra this buffer struct should probably simply be a single sync map
		// for all prepared buffers, e.g. energy, player, etc. We have to merge all
		// fanout buffers into one single slice to simplify the reconciliation below
		// in Engine.worker().

		var b [][]byte
		{
			b, _ = e.buf.ply.LoadAndDelete(k)
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
	}
}

func (e *Engine) worker(ply *player.Player, buf [][]byte, tim *time.Timer) {
	// The semaphore below controls the amount of workers that are allowed to
	// process packets at the same time. Every time we process a player, we
	// allocate a ticket with the semaphore, before doing the actual work.

	{
		e.sem <- struct{}{}
	}

	// Use a dedicated channel for signalling different outcomes. Closing either
	// of the channels below does not conflict in any way with the others. We use
	// the early channel to stop processing before deadline violation. And we use
	// the late channel in order to cleanup the goroutine after our deadline has
	// expired.

	var ear chan struct{}
	var lat chan struct{}
	{
		ear = make(chan struct{})
		lat = make(chan struct{})
	}

	// Send the prepared fanout buffers to the given player. The use of the early
	// channel below allows us to proceed before our deadline expires. The use of
	// the late channel below allows is to prevent any more work once the given
	// deadline gets violated.

	go func() {
		for _, x := range buf {
			select {
			case <-lat:
				return
			default:
				ply.Cli.Stream(x)
			}
		}

		// In case we processed all fanout buffers in time, close the early channel
		// so we can release the allocated semaphore. The early channel will not be
		// closed if streaming above blocks on any but the very last buffer, because
		// if a blocked stream gets unblocked, and if the for loop executes another
		// iteration while the late channel got closed, the goroutine returns early.
		// If on the other hand the streaming of the very last buffer is blocked,
		// then the for loop will not check the select anymore, but instead break
		// and cause the early channel to be closed.

		{
			close(ear)
		}
	}()

	// Block this worker call until we are done processing the given player, or
	// until the provided deadline expired.

	select {
	case <-ear:
		tim.Stop()
	case <-tim.C:
		close(lat)
	}

	// Ensure to release the semaphore that we allocated for this worker call.

	{
		<-e.sem
	}
}
