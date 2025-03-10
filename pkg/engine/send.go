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
		// Skip all inactive players.

		if !e.act[u] {
			continue
		}

		// Forward the fanout buffer to the client specific goroutine for capacity
		// aware processing. The buffer channels provided by each client must never
		// block.

		{
			e.fcn[u] <- e.fbf[u]
		}

		// Reset the player specific fanout buffer for the next cycle.

		{
			e.fbf[u] = nil
		}
	}
}
