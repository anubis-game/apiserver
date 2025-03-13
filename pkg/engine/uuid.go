package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/ethereum/go-ethereum/common"
)

func (e *Engine) uuid(pac router.Uuid) {
	if pac.Jod == router.Join {
		e.join(pac.Uid, pac.Wal, pac.Cli)
	} else {
		e.drop(pac.Uid)
	}
}

func (e *Engine) join(u byte, w common.Address, c chan<- []byte) {
	// Generating a new player object for the connected client effectively puts
	// the player randomly onto the game map due to the Filler.Vector()
	// randomization.

	var v *vector.Vector
	{
		v = e.fil.Vector()
	}

	var f []byte

	{
		f = append(f, byte(schema.Uuid), u)
		f = append(f, w.Bytes()...)
		f = append(f, byte(schema.Body), u, 0x0)
	}

	l := len(f) - 1
	v.Ranger(func(c matrix.Coordinate) {
		// Add the new byte ID to the partition indices.

		{
			e.lkp.add(u, c)
		}

		// Add the new Vector's node coordinates to the new fanout buffer.

		{
			f[l]++
			a := c.Byt()
			f = append(f, a[:]...)
		}
	})

	// Search for all the energy packets located within the partitions that the
	// new player can see.

	for _, p := range v.Layers(v.Charax().Fos, matrix.Pt1) {
		for k := range e.lkp.nrg[p] {
			f = append(f, e.mem.nrg[k]...)
		}
	}

	// Render all existing players inside the view of the new player, and render
	// the new player in the view of all existing players.

	l, m, n, o := v.Bounds()

	e.screen(u, v.Change().Hea.Pt8(), func(b byte, w *vector.Vector) {
		w.Inside(l, m, n, o, func(c matrix.Coordinate) bool {
			b := c.Byt()
			// TODO:infra the body messages still need to be encoded.
			f = append(f, b[:]...)
			return true
		})

		p, q, r, s := w.Bounds()

		v.Inside(p, q, r, s, func(c matrix.Coordinate) bool {
			b := c.Byt()
			// TODO:infra the body messages still need to be encoded.
			f = append(f, b[:]...)
			return true
		})
	})

	// TODO:infra send the new player information to all players that can see the
	// new Vector, but make sure to not update fanout buffers of active
	// disconnected players.

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward. Also store the player's
	// buffer in the player's setter.

	{
		e.ply.act[u] = true
		e.ply.qdr[u] = v.Motion().Qdr
		e.ply.buf[u] = f
		e.ply.cli[u] = c
		e.ply.agl[u] = v.Motion().Agl
		e.mem.vec[u] = v
	}
}

func (e *Engine) drop(u byte) {
	e.ply.qdr[u] = 0
	e.ply.buf[u] = nil
	e.ply.cli[u] = nil // TODO:test ensure we can concurrently read while a single writer modifies the fanout channel
	e.ply.agl[u] = 0
	e.mem.vec[u] = nil
}
