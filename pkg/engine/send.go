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

		// Forward the fanout buffer to the client specific goroutine for capacity
		// aware processing. The buffer channels provided by each client must never
		// block. Client specific fanout channels may be nil if active players
		// disconnected.

		if cli != nil {
			cli <- e.ply.buf[u]
		}

		// Reset the player specific fanout buffer for the next cycle. We keep the
		// existing sequence byte and prevent re-allocation of the underlying data
		// array using the resclicing operation [:1]. In case active players have no
		// connected client, we discard all fanout buffers without sending, until
		// the player comes back online or dies.

		{
			e.ply.buf[u] = e.ply.buf[u][:1]
		}

		// Increment the sequence byte for the next cycle.

		{
			e.ply.buf[u][0]++
		}
	}
}
