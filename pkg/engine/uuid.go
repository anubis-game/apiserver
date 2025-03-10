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

func (e *Engine) join(uid byte, _ common.Address, cli chan<- []byte) {
	all := map[byte]struct{}{}

	// Generating a new player object for the connected client effectively puts
	// the player randomly onto the game map due to the Filler.Vector()
	// randomization.

	var vec *vector.Vector
	{
		vec = e.fil.Vector()
	}

	// We separate the player identification from the vector representation. The
	// body parts are associated with the player ID the same way the user's wallet
	// is associated with that same player ID.

	var ini []byte
	{
		ini = make([]byte, 61) // 22 + 33 + 3 + 3

		// copy(ini[:22], ply.Wallet())   // len(22)
		copy(ini[22:55], vec.Encode()) // len(33)
		// copy(ini[55:58], crx.Size())   // len(3)
		// copy(ini[58:61], crx.Type())   // len(3)
	}

	// Send the new player's own wallet information first so every player can self
	// identify. Also send the players own body parts and motion configuration.

	var buf []byte
	{
		buf = ini
	}

	// Search for all the energy packets located within the partitions that the
	// new player can see.

	for _, p := range matrix.Pt1Scr(vec.Screen()) {
		for k := range e.lkp.nrg[p] {
			buf = append(buf, e.mem.nrg[k].Encode()...)
		}
	}

	// Search for all the unique byte IDs located around the new player's head
	// node. The area we are searching through here contains a single layer of
	// large partitions around the new player's head node. That is 9 partitions.

	for _, p := range matrix.Pt8Scr(vec.Header().Pt8()) {
		for u := range e.lkp.pt8[p] {
			all[u] = struct{}{}
		}
	}

	// Render all existing players inside the view of the new player, and render
	// the new player in the view of all existing players.

	for u := range all {
		var v *vector.Vector
		{
			v = e.mem.vec[u]
		}

		for _, l := range v.Inside(vec.Screen()) {
			for _, c := range l {
				b := c.Byt()
				// TODO:infra the body messages still need to be encoded.
				buf = append(buf, b[:]...)
			}
		}

		for _, l := range vec.Inside(v.Screen()) {
			for _, c := range l {
				b := c.Byt()
				// TODO:infra the body messages still need to be encoded.
				buf = append(buf, b[:]...)
			}
		}
	}

	// Add the new byte ID to the partition indices.

	vec.Ranger(func(c matrix.Coordinate) {
		// Add every node coordinate of the new player to the small partitions.

		{
			prt := c.Pt1()

			pt1, exi := e.lkp.pt1[prt]
			if !exi {
				pt1 = map[byte]struct{}{uid: {}}
			} else {
				pt1[uid] = struct{}{}
			}

			e.lkp.pt1[prt] = pt1
		}

		// Add every node coordinate of the new player to the large partitions.

		{
			prt := c.Pt8()

			pt8, exi := e.lkp.pt8[prt]
			if !exi {
				pt8 = map[byte]struct{}{uid: {}}
			} else {
				pt8[uid] = struct{}{}
			}

			e.lkp.pt8[prt] = pt8
		}
	})

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward. Also store the player's
	// buffer in the player's setter.

	{
		e.ply.buf[uid] = buf
		e.ply.cli[uid] = cli
		e.mem.vec[uid] = vec
	}
}

func (e *Engine) drop(uid byte) {
	e.ply.buf[uid] = nil
	e.ply.cli[uid] = nil // TODO:test ensure we can concurrently read while a single writer modifies the fanout channel
	e.mem.vec[uid] = nil
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
