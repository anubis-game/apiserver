package engine

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/vector"
)

func (e *Engine) uuid(pac router.Packet) {
	// Generating a new player object for the connected client effectively puts
	// the player randomly onto the game map due to the Filler.Vector()
	// randomization.

	var vec *vector.Vector
	{
		vec = e.fil.Vector(pac.Uid)
	}

	var crx vector.Charax
	{
		crx = vec.Charax().Get()
	}

	var ply *player.Player
	{
		ply = player.New(player.Config{
			Cli: pac.Cli,
			Uid: pac.Uid,
			Vec: vec,
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
		copy(ini[55:58], crx.Size())   // len(3)
		copy(ini[58:61], crx.Type())   // len(3)
	}

	// Send the new player's own wallet information first so every player can self
	// identify. Also send the players own body parts and motion configuration.

	var buf []byte
	{
		buf = ini
	}

	e.mem.ply.Range(func(k byte, v *player.Player) bool {
		// Only add the new player's body parts to the current view of an existing
		// player, if the body parts of the new player are visible inside the view
		// of the existing player. Note that every player joining the game must push
		// its own identity to all active players first, so the following body parts
		// can be identified using the player's byte Id. Further note that any
		// buffer modifications of existing players must be synchronized, which is
		// why we are using MapOf.Compute() below.

		if vec.Inside(v.Vec.Screen()) {
			e.buf.Compute(pac.Uid, func(b []byte, _ bool) ([]byte, bool) {
				b = append(b, ini...)
				return b, false
			})
		}

		// Only share the existing player's information with the new player, if the
		// body parts of the existing player are visible inside the view of the new
		// player.

		if v.Vec.Inside(vec.Screen()) {
			var c vector.Charax
			{
				c = v.Vec.Charax().Get()
			}

			{
				buf = append(buf, v.Wallet()...)
				buf = append(buf, c.Size()...)
				buf = append(buf, c.Type()...) // TODO:infra how can we prevent sending type twice?
			}
		}

		return true
	})

	// Stream all relevant map details for the initial view of the new player.
	// This process implies to find all relevant energy and player details visible
	// to the new player.

	for _, x := range vec.Screen().Prt {
		{
			// Search for all the energy packets located within the partition x.

			var lkp map[object.Object]struct{}
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
			// Search for all the player segments located within the partition x.

			lkp, _ := e.lkp.ply.Load(x)

			// For every player segment in partition x, add its encoded representation to
			// the new player's fanout buffer.

			for k := range lkp {
				var p *player.Player
				{
					p, _ = e.mem.ply.Load(k)
				}

				{
					buf = append(buf, p.Vec.Buffer(x)...)
				}
			}
		}
	}

	// Add the new player to the lookup table based on its currently occupied
	// coordinates.

	for _, x := range vec.Occupy().Prt {
		e.lkp.ply.Compute(x, func(old map[byte]struct{}, exi bool) (map[byte]struct{}, bool) {
			if !exi {
				old = map[byte]struct{}{}
			}

			{
				old[pac.Uid] = struct{}{}
			}

			return old, false
		})
	}

	// After we add all initially occupied partitions to the lookup table, we can
	// reset the occupied partitions once below, because all further updates on
	// the Vector's occupied partitions will be tracked using the Occupy.New and
	// Occupy.Old fields. The reason for this is the fact that once initialized, a
	// Vector does only ever change one occupied partition at a time.

	{
		vec.Occupy().Prt = nil
	}

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward. Also store the player's
	// buffer in the player's setter.

	{
		e.buf.Store(pac.Uid, buf)
		e.mem.ply.Store(pac.Uid, ply)
	}
}
