package engine

import (
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
)

func (e *Engine) join(pac router.Packet) {
	// Generating a new player object for the connected client effectively puts
	// the player randomly onto the game map due to the Filler.Vector()
	// randomization.

	var ply *player.Player
	{
		ply = &player.Player{
			Cli: pac.Cli,
			Uid: pac.Uid,
			Vec: e.fil.Vector(),
		}
	}

	// We separate the player identification from the vector representation. The
	// body parts are associated with the 2 byte player ID the same way the user's
	// wallet is associated with that same 2 byte player ID.

	var bod []byte
	var joi []byte
	{
		bod = schema.Encode(schema.Body, ply.EncodeVector())
		joi = schema.Encode(schema.Join, ply.EncodeWallet())
	}

	// Send the new player's own ID first so every player can self identify.

	e.buf.ply.Compute(ply.Uid, func(old [][]byte, _ bool) ([][]byte, bool) {
		return append(old, joi), false
	})

	for k, v := range e.mem.ply {
		// Only add the fanout buffer to the current view of an existing player, if
		// the body of the new player is visible inside the view of the existing
		// player.

		if ply.Vec.Inside(v.Vec.Screen()) {
			e.buf.ply.Compute(k, func(old [][]byte, _ bool) ([][]byte, bool) {
				return append(old, bod), false
			})
		}

		// Every player joining the game must push its own identity to all active
		// players.

		e.buf.ply.Compute(k, func(old [][]byte, _ bool) ([][]byte, bool) {
			return append(old, joi), false
		})

		// Every player joining a game must receive the full list of active players,
		// so that we can associate a player's 2 byte IDs with their respective 20
		// byte wallets.

		e.buf.ply.Compute(ply.Uid, func(old [][]byte, _ bool) ([][]byte, bool) {
			return append(old, schema.Encode(schema.Join, v.EncodeWallet())), false
		})
	}

	// Stream all relevant map details for the initial view of the new player.
	// This process implies to find all relevant energy and player details visible
	// to the new player.

	for _, x := range ply.Vec.Bounds() {
		{
			// Search for all the energy packets located within the partition x.

			lkp, _ := e.lkp.nrg.Load(x)

			// For every energy packet in partition x, add its encoded representation to
			// the new player's fanout buffer.

			for k := range lkp {
				e.buf.ply.Compute(ply.Uid, func(old [][]byte, _ bool) ([][]byte, bool) {
					return append(old, schema.Encode(schema.Food, e.mem.nrg[k].Encode())), false
				})
			}
		}

		{
			// Search for all the player segments located within the partition x.

			lkp, _ := e.lkp.ply.Load(x)

			// For every player segment in partition x, add its encoded representation to
			// the new player's fanout buffer.

			for k := range lkp {
				e.buf.ply.Compute(ply.Uid, func(old [][]byte, _ bool) ([][]byte, bool) {
					return append(old, schema.Encode(schema.Body, e.mem.ply[k].EncodeBuffer(x))), false
				})
			}
		}
	}

	// Add the new player to the lookup table based on its currently occupied
	// coordinates.

	for _, x := range ply.Vec.Occupy() {
		e.lkp.ply.Compute(x, func(old map[[2]byte]struct{}, exi bool) (map[[2]byte]struct{}, bool) {
			if !exi {
				old = map[[2]byte]struct{}{}
			}

			{
				old[ply.Uid] = struct{}{}
			}

			return old, false
		})
	}

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward.

	{
		e.mem.ply[pac.Uid] = ply
	}
}
