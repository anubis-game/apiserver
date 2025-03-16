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
		e.join(pac.Uid, pac.Wal, pac.Cli, pac.Vec)
	} else {
		e.drop(pac.Uid)
	}
}

func (e *Engine) join(u byte, a common.Address, c chan<- []byte, v *vector.Vector) {
	var f []byte

	{
		f = append(f, byte(schema.Uuid), u)      // 2
		f = append(f, a.Bytes()...)              // 20
		f = append(f, byte(schema.Body), u, 0x0) // 3
	}

	v.Ranger(func(c matrix.Coordinate) {
		// Add the new byte ID to the partition indices.

		{
			e.lkp.add(u, c)
		}

		// Add the new Vector's node coordinates to the new fanout buffer.

		{
			f[24]++ // 25 - 1
			x := c.Byt()
			f = append(f, x[:]...)
		}
	})

	// Search for all the energy packets located within the partitions that the
	// new player can see.

	e.energy(v, func(e []byte) {
		f = append(f, e...) // TODO:infra the energy bytes need to be schema encoded
	})

	// Render all existing players inside the view of the new player, and render
	// the new player in the view of all existing players.

	var t, r, b, l int
	{
		t, r, b, l = v.Screen(matrix.Pt1)
	}

	e.screen(v, func(w *vector.Vector) {
		w.Inside(t, r, b, l, func(c matrix.Coordinate) bool {
			x := c.Byt()
			f = append(f, x[:]...) // TODO:infra the body messages still need to be encoded.
			return true
		})

		g, h, i, j := w.Screen(matrix.Pt1)

		v.Inside(g, h, i, j, func(c matrix.Coordinate) bool {
			x := c.Byt()
			f = append(f, x[:]...) // TODO:infra the body messages still need to be encoded.
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
