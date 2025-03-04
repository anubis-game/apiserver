package engine

import (
	"time"
)

func (e *Engine) send(tic time.Time) {
	// TODO:metrics monitor ( tic - e.tic ) to see how regular our fanout
	// procedure executes throughout the program lifetime.

	{
		e.tic = tic
	}

	// Distribute the work across all client specific goroutines. This part must
	// be called sequentially, because we reset every player's fanout buffer. If
	// we were to run the iterations below concurrently, then we would potentially
	// delete the fanout buffer of a goroutine that has not started to write the
	// data that it meant to send out to the client.

	for u := range e.uni.Length() {
		// Get the player specific buffer and reset it for the next cycle. Buffers
		// may be empty if player IDs have been allocated upon joining a game, while
		// no buffer has been prepared just yet.

		var b []byte
		{
			b, _ = e.fbf.LoadAndDelete(u)
		}

		if len(b) == 0 {
			continue
		}

		// Get the player specific fanout channel so we can forward the prepared
		// fanout buffer to the underlying client. Channels may be nil if player IDs
		// have been allocated upon joining a game, while no client has been setup
		// just yet. It is also possible for players to get disconnected
		// intermittently, which would nil the formerly established channel as well.

		var c chan<- []byte
		{
			c = e.fcn[u]
		}

		if c == nil {
			continue
		}

		// Forward the fanout buffer to the client specific goroutine for capacity
		// aware processing. The buffer channels provided by each client must never
		// block.

		{
			c <- b
		}
	}
}
