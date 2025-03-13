package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/router"
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

func (e *Engine) join(u byte, _ common.Address, c chan<- []byte) {
	// Generating a new player object for the connected client effectively puts
	// the player randomly onto the game map due to the Filler.Vector()
	// randomization.

	var v *vector.Vector
	{
		v = e.fil.Vector()
	}

	// We separate the player identification from the vector representation. The
	// body parts are associated with the player ID the same way the user's wallet
	// is associated with that same player ID.

	var ini []byte
	{
		ini = make([]byte, 61) // 22 + 33 + 3 + 3

		// copy(ini[:22], ply.Wallet())   // len(22)
		copy(ini[22:55], v.Encode()) // len(33)
		// copy(ini[55:58], crx.Size())   // len(3)
		// copy(ini[58:61], crx.Type())   // len(3)
	}

	// TODO:infra send the new Vector's own body to themselves

	// Send the new player's own wallet information first so every player can self
	// identify. Also send the players own body parts and motion configuration.

	var f []byte
	{
		f = ini
	}

	// Search for all the energy packets located within the partitions that the
	// new player can see.

	for _, p := range v.Layers(v.Charax().Fos, matrix.Pt1) {
		for k := range e.lkp.nrg[p] {
			f = append(f, e.mem.nrg[k].Encode()...)
		}
	}

	// Render all existing players inside the view of the new player, and render
	// the new player in the view of all existing players.

	l, m, n, o := v.Bounds()

	for _, b := range e.allpt8(u, v) {
		var w *vector.Vector
		{
			w = e.mem.vec[b]
		}

		for _, c := range w.Inside(l, m, n, o) {
			b := c.Byt()
			// TODO:infra the body messages still need to be encoded.
			f = append(f, b[:]...)
		}

		for _, c := range v.Inside(w.Bounds()) {
			b := c.Byt()
			// TODO:infra the body messages still need to be encoded.
			f = append(f, b[:]...)
		}
	}

	// Add the new byte ID to the partition indices.

	v.Ranger(func(crd matrix.Coordinate) {
		e.lkp.add(u, crd)
	})

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

// var uid []byte
// {
// 	uid = make([]byte, 22)
// }
//
// {
// 	uid[0] = byte(schema.Uuid)
// 	uid[1] = c.Uid
// }
//
// {
// 	copy(uid[2:], c.Wal.Bytes())
// }
