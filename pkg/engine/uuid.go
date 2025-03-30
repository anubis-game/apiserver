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
		fa = make([]byte, 25)
	}

	{
		b := [2]byte{byte(schema.Uuid), ua}
		copy(fa[:2], b[:])
	}

	{
		copy(fa[2:22], wa[:])
	}

	{
		b := [2]byte{byte(schema.Body), ua}
		copy(fa[22:24], b[:])
	}

	va.Ranger(func(c matrix.Coordinate) {
		// Add the new Vector's node coordinates to the new fanout buffer.

		{
			fa[24]++
			b := c.Byt()
			fa = append(fa, b[:]...)
		}

		// Add the new byte ID to the partition indices.

		{
			e.lkp.add(ua, c)
		}
	})

	// Search for all the energy packets located within the partitions that the
	// new player can see. Note that the stored energy slices are already schema
	// encoded.

	{
		fa = append(fa, byte(schema.Food), 0x0)
	}

	e.energy(va, func(e []byte) {
		fa[38]++
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
	e.ply.cli[u] = nil
	e.ply.agl[u] = 0
	e.mem.vec[u] = nil
}
