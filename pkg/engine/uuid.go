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

func (e *Engine) join(ua byte, wa common.Address, ca chan<- []byte, va *vector.Vector) {
	var fa []byte

	{
		fa = append(fa, byte(schema.Uuid), ua)      // 2
		fa = append(fa, wa.Bytes()...)              // 20
		fa = append(fa, byte(schema.Body), ua, 0x0) // 3
	}

	va.Ranger(func(c matrix.Coordinate) {
		// Add the new byte ID to the partition indices.

		{
			e.lkp.add(ua, c)
		}

		// Add the new Vector's node coordinates to the new fanout buffer.

		{
			fa[24]++ // 25 - 1
			x := c.Byt()
			fa = append(fa, x[:]...)
		}
	})

	// Search for all the energy packets located within the partitions that the
	// new player can see. Note that the stored energy slices are already schema
	// encoded.

	e.energy(va, func(e []byte) {
		fa = append(fa, e...)
	})

	// Render all existing players inside the view of the new player, and render
	// the new player inside the view of all existing players, but make sure to
	// not update fanout buffers of active disconnected players.

	ta, ra, ba, la := va.Screen(matrix.Pt1)
	e.screen(va, func(ub byte, vb *vector.Vector) {
		if e.ply.cli[ub] != nil {
			tb, rb, bb, lb := vb.Screen(matrix.Pt1)
			e.ply.buf[ub] = e.inside(e.ply.buf[ub], va, ua, tb, rb, bb, lb)
		}

		{
			fa = e.inside(fa, vb, ub, ta, ra, ba, la)
		}
	})

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward. Also store the player's
	// buffer in the player's setter.

	{
		e.ply.act[ua] = true
		e.ply.qdr[ua] = va.Motion().Qdr
		e.ply.buf[ua] = fa
		e.ply.cli[ua] = ca
		e.ply.agl[ua] = va.Motion().Agl
		e.mem.vec[ua] = va
	}
}

func (e *Engine) drop(u byte) {
	e.ply.qdr[u] = 0
	e.ply.buf[u] = nil
	e.ply.cli[u] = nil // TODO:test ensure we can concurrently read while a single writer modifies the fanout channel
	e.ply.agl[u] = 0
	e.mem.vec[u] = nil
}
