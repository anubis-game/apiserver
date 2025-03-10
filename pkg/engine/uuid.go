package engine

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/ethereum/go-ethereum/common"
)

func (e *Engine) uuid(pac router.Uuid) {
	if pac.Jod == router.Join {
		e.join(pac.Uid, pac.Wal, pac.Fcn)
	} else {
		e.drop(pac.Uid)
	}
}

func (e *Engine) join(uid byte, _ common.Address, fcn chan<- []byte) {

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

	// Stream all relevant map details for the initial view of the new player.

	all := map[byte]struct{}{}
	for _, x := range matrix.PfromS(vec.Screen()) {
		{
			// Search for all the energy packets located within the partition x.

			var lkp map[matrix.Coordinate]struct{}
			{
				lkp, _ = e.lkp.nrg.Load(x)
			}

			// For every energy packet in partition x, add its encoded representation to
			// the new player's fanout buffer.

			for k := range lkp {
				var n *energy.Energy
				{
					n, _ = e.mem.nrg.Load(k)
				}

				{
					buf = append(buf, n.Encode()...)
				}
			}
		}

		{
			// Search for all the unique byte IDs located within partition x.

			var lkp map[byte]struct{}
			{
				lkp, _ = e.lkp.ply.Load(x)
			}

			for k := range lkp {
				all[k] = struct{}{}
			}
		}
	}

	// Render all existing players inside the view of the new player.

	for u := range all {
		var v *vector.Vector
		{
			v, _ = e.mem.vec.Load(u)
		}

		for _, l := range v.Inside(vec.Screen()) {
			for _, c := range l {
				b := c.Byt()
				// TODO:infra the body messages still need to be encoded.
				buf = append(buf, b[:]...)
			}
		}
	}

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward. Also store the player's
	// buffer in the player's setter.

	{
		e.fbf[uid] = buf
		e.fcn[uid] = fcn
		e.mem.vec.Store(uid, vec)
	}

	// As the very last step, activate the joined player for further our fanout
	// procedure.

	{
		e.act[uid] = true
	}
}

func (e *Engine) drop(uid byte) {
	e.act[uid] = false
	e.fbf[uid] = nil
	e.fcn[uid] = nil // TODO:test ensure we can concurrently read while a single writer modifies the fanout channel
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
