package engine

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/google/uuid"
)

func (e *Engine) fanout(tic time.Time) {
	{
		// TODO monitor tic - e.tic
	}

	{
		e.tic = tic
	}

	for k, v := range e.mem.ply {
		var n [][]byte
		var p [][]byte
		{
			n, _ = e.buf.nrg.Load(k)
			p, _ = e.buf.ply.Load(k)
		}

		{
			go e.worker(k, v, n, p)
		}
	}
}

func (e *Engine) worker(_ uuid.UUID, ply *player.Player, nbf [][]byte, pbf [][]byte) {
	// The semaphore controls the amount of workers that are allowed to process
	// packets at the same time. Every time we receive a packet, we push a ticket
	// into the semaphore before doing the work.
	{
		e.sem <- struct{}{}
	}

	// Send player movements of the enemies first so that every player can
	// react based on the full picture of the current frame.

	// TODO add prepared player bytes to wallet address, if any
	for _, x := range pbf {
		ply.Cli.Stream(schema.Encode(schema.Move, x))
	}

	// TODO check for wallet specific movement and calculate Target(). We
	// cannot just fanout a prepared byte slice here, since we have to force
	// the player movement in either the currently chosen, or the latest known
	// direction.
	// cli.Window()

	// Send energy changes last, since player updates are more relevant.

	// TODO add prepared energy bytes to wallet address, if any
	for _, x := range nbf {
		ply.Cli.Stream(schema.Encode(schema.Food, x)) // TODO this action encoding should not happen here
	}

	// Ensure we remove our ticket from the semaphore once all work was completed.
	{
		<-e.sem
	}
}
