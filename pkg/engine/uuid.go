package engine

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
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
		vec = e.fil.Vector(uid)
	}

	var crx vector.Charax
	{
		crx = vec.Charax().Get()
	}

	var fuw []byte
	{
		fuw = make([]byte, 22)

		fuw[0] = byte(schema.Uuid)
		fuw[1] = uid

		copy(fuw[2:], wal.Bytes())
	}

	// We separate the player identification from the vector representation. The
	// body parts are associated with the player ID the same way the user's wallet
	// is associated with that same player ID.

	var ini []byte
	{
		ini = make([]byte, 61) // 22 + 33 + 3 + 3

		copy(ini[:22], fuw)            // len(22)
		copy(ini[22:55], vec.Encode()) // len(33)
		copy(ini[55:58], crx.Size())   // len(3)
		copy(ini[58:61], crx.Type())   // len(3)
	}

	// Send the new player's own wallet information first so every player can self
	// identify. Also send the players own body parts and motion configuration.

	var buf []byte
	{
		buf = ini
	}

	// Iterate over all allocated Vectors. This loop will always iterate over
	// Config.Cap items.

	for u := range e.uni.Length() {
		// Skip all unallocated Vectors.

		var v *vector.Vector
		{
			v = e.mvc[u].Load()
		}

		if v == nil {
			continue
		}

		// Only add the new player's body parts to the current view of an existing
		// player, if the body parts of the new player are visible inside the view
		// of the existing player. Note that every player joining the game must push
		// its own identity to all active players first, so the following body parts
		// can be identified using the player's byte Id. Further note that any
		// buffer modifications of existing players must be synchronized, which is
		// why we are using MapOf.Compute() below.

		if vec.Inside(v.Screen()) {
			//
		}

		// Only share the existing player's information with the new player, if the
		// body parts of the existing player are visible inside the view of the new
		// player.

		if v.Inside(vec.Screen()) {
			var c vector.Charax
			{
				c = v.Charax().Get()
			}

			{
				buf = append(buf, e.fuw[u]...)
				buf = append(buf, c.Size()...)
				buf = append(buf, c.Type()...) // TODO:infra how can we prevent sending type twice?
			}
		}
	}

	// Stream all relevant map details for the initial view of the new player.
	// This process implies to find all relevant energy and player details visible
	// to the new player.

	for _, x := range vec.Screen().Prt {
		{
			// Search for all the energy packets located within the partition x.

			var lkp map[object.Object]struct{}
			{
				lkp, _ = e.pen.Load(x)
			}

			// For every energy package in partition x, add its encoded representation
			// to the new player's fanout buffer.

			for o := range lkp {
				var n *energy.Energy
				{
					n, _ = e.men.Load(o)
				}

				{
					buf = append(buf, n.Encode()...)
				}
			}
		}

		{
			// Search for all the player segments located within the partition x.

			lkp, _ := e.pvc.Load(x)

			// For every player segment in partition x, add its encoded representation to
			// the new player's fanout buffer.

			for u := range lkp {
				var v *vector.Vector
				{
					v = e.mvc[u].Load()
				}

				{
					buf = append(buf, v.Buffer(x)...)
				}
			}
		}
	}

	// Add the new player to the lookup table based on its currently occupied
	// coordinates.

	for _, x := range vec.Occupy().Prt {
		e.pvc.Compute(x, func(old map[byte]struct{}, exi bool) (map[byte]struct{}, bool) {
			if !exi {
				old = map[byte]struct{}{}
			}

			{
				old[uid] = struct{}{}
			}

			return old, false
		})
	}

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward. Also store the player's
	// buffer in the player's setter.

	{
		e.fbf[uid].Store(&buf)
		e.fuw[uid] = fuw
		e.mvc[uid].Store(vec)
	}

	// As the very last step, assign the player's fanout channel
	{
		e.fcn[uid] = fcn
	}
}

func (e *Engine) drop(uid byte) {
	e.fcn[uid] = nil // TODO:test ensure we can concurrently read while a single writer modifies the fanout channel
}
