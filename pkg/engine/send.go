package engine

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/player"
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
			b, _ = e.buf.LoadAndDelete(u)
		}

		if len(b) == 0 {
			continue
		}

		// Get the player specific client so we can stream the prepared fanout
		// buffer.  Player's may be nil if player IDs have been allocated upon
		// joining a game, while no client has been setup just yet.

		// TODO:infra given that Engine.uuid() is the only sequential writer for the
		// player clients, we can use a simple slice indexed by the player byte IDs
		// in order to store and read the client pointers.

		var p *player.Player
		{
			p, _ = e.mem.ply.Load(u)
		}

		if p == nil {
			continue
		}

		// Forward the fanout buffer to the client specific goroutine for capacity
		// aware processing. The buffer channels provided by each client must never
		// block.

		// TODO:infra we do not want to remove players from the game if their client
		// connection is accidentally flaky. But a disconnected client cannot
		// consume messages anymore, causing the fanout buffers below to fill up.
		// Such congested buffer channels block the entire fanout procedure, which
		// must remain linear up to this point. If we want to keep disconnected
		// players in the game, then we have to stop serving them below until they
		// can process messages again.
		//
		//     1. What is the mechanism to decide whether to send any more messages
		//        to any given client?
		//
		//     2. How do we restore a websocket connection for players that never
		//        left the game?
		//

		{
			p.Cli.Buffer() <- b
		}
	}
}
