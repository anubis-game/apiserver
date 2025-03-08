package engine

import (
	"github.com/anubis-game/apiserver/pkg/player"
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

func (e *Engine) join(uid byte, wal common.Address, fcn chan<- []byte) {

	// Generating a new player object for the connected client effectively puts
	// the player randomly onto the game map due to the Filler.Vector()
	// randomization.

	var vec *vector.Vector
	{
		vec = e.fil.Vector()
	}

	var ply *player.Player
	{
		ply = player.New(player.Config{
			Uid: uid,
			Vec: vec,
			Wal: wal,
		})
	}

	// We separate the player identification from the vector representation. The
	// body parts are associated with the player ID the same way the user's wallet
	// is associated with that same player ID.

	var ini []byte
	{
		ini = make([]byte, 61) // 22 + 33 + 3 + 3

		copy(ini[:22], ply.Wallet())   // len(22)
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
	// This process implies to find all relevant energy and player details visible
	// to the new player.

	// for _, x := range vec.Screen() {
	// 	{
	// 		// Search for all the energy packets located within the partition x.

	// 		var lkp map[matrix.Coordinate]struct{}
	// 		{
	// 			lkp, _ = e.lkp.nrg.Load(x)
	// 		}

	// 		// For every energy packet in partition x, add its encoded representation to
	// 		// the new player's fanout buffer.

	// 		for k := range lkp {
	// 			var n *energy.Energy
	// 			{
	// 				n, _ = e.mem.nrg.Load(k)
	// 			}

	// 			{
	// 				buf = append(buf, n.Encode()...)
	// 			}
	// 		}
	// 	}

	// 	{
	// 		// Search for all the player coordinates located within the partition x.

	// 		lkp, _ := e.lkp.ply.Load(x)

	// 		// For every player coordinate in partition x, add its encoded
	// 		// representation to the new player's fanout buffer.

	// 		for k := range lkp {
	// 			var p *player.Player
	// 			{
	// 				p, _ = e.mem.ply.Load(k)
	// 			}

	// 			for _, y := range p.Vec.Ocdiff(x) {
	// 				b := y.Byt()
	// 				buf = append(buf, b[:]...)
	// 			}
	// 		}
	// 	}
	// }

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward. Also store the player's
	// buffer in the player's setter.

	{
		e.fbf.Store(uid, buf)
		e.mem.ply.Store(uid, ply)
	}

	// As the very last step, assign the player's fanout channel
	{
		e.fcn[uid] = fcn
	}
}

func (e *Engine) drop(uid byte) {
	e.fcn[uid] = nil // TODO:test ensure we can concurrently read while a single writer modifies the fanout channel
}
