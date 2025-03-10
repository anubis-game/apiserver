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
		var cli chan<- []byte
		{
			cli = e.ply.cli[u]
		}

		// Skip all inactive players.

		if cli == nil {
			continue
		}

		// Forward the fanout buffer to the client specific goroutine for capacity
		// aware processing. The buffer channels provided by each client must never
		// block.

		{
			cli <- e.ply.buf[u]
		}

		// Reset the player specific fanout buffer for the next cycle.

		{
			e.ply.buf[u] = nil
		}
	}
}
