package engine

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/google/uuid"
)

func (e *Engine) push(tic time.Time) {
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

		// TODO any client becoming a liability by blocking available semaphores for
		// too long have to be terminated in order to address any downstream
		// bottlenecks.

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

	for _, x := range pbf {
		ply.Cli.Stream(x)
	}

	// Send energy changes last, since player updates are more relevant.

	for _, x := range nbf {
		ply.Cli.Stream(x)
	}

	// Ensure we remove our ticket from the semaphore once all work was completed.
	{
		<-e.sem
	}
}
