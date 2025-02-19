package engine

import (
	"runtime"
	"time"

	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/vector"
)

const (
	// Cap is the duration based capacity that we allow for a single fanout
	// procedure to take per worker process. E.g. a standard frame duration of 25
	// milliseconds implies a total amount of 24 milliseconds per fanout procedure
	// per worker, given an overhead buffer of 1 millisecond that we may incur at
	// runtime each cycle.
	Cap = time.Duration(vector.Frm-1) * time.Millisecond
)

func (e *Engine) join(pac router.Packet) {
	// Generating a new player object for the connected client effectively puts
	// the player randomly onto the game map due to the Filler.Vector()
	// randomization.

	var ply *player.Player
	{
		ply = player.New(player.Config{
			Cli: pac.Cli,
			Uid: pac.Uid,
			Vec: e.fil.Vector(pac.Uid),
		})
	}

	// We separate the player identification from the vector representation. The
	// body parts are associated with the 2 byte player ID the same way the user's
	// wallet is associated with that same 2 byte player ID.

	var bod []byte
	var joi []byte
	{
		bod = schema.Encode(schema.Body, ply.Vector())
		joi = schema.Encode(schema.Join, ply.Wallet())
	}

	// Send the new player's own ID first so every player can self identify.

	e.buf.Compute(pac.Uid, func(old []byte, _ bool) ([]byte, bool) {
		return append(old, joi...), false
	})

	e.mem.ply.Range(func(k [2]byte, v *player.Player) bool {
		// Only add the fanout buffer to the current view of an existing player, if
		// the body of the new player is visible inside the view of the existing
		// player.

		if ply.Vec.Inside(v.Vec.Screen()) {
			e.buf.Compute(k, func(old []byte, _ bool) ([]byte, bool) {
				return append(old, bod...), false
			})
		}

		// Every player joining the game must push its own identity to all active
		// players.

		e.buf.Compute(k, func(old []byte, _ bool) ([]byte, bool) {
			return append(old, joi...), false
		})

		// Every player joining a game must receive the full list of active players,
		// so that we can associate a player's 2 byte IDs with their respective 20
		// byte wallets.

		e.buf.Compute(pac.Uid, func(old []byte, _ bool) ([]byte, bool) {
			return append(old, schema.Encode(schema.Join, v.Wallet())...), false
		})

		return true
	})

	// Stream all relevant map details for the initial view of the new player.
	// This process implies to find all relevant energy and player details visible
	// to the new player.

	for _, x := range ply.Vec.Screen().Prt {
		{
			// Search for all the energy packets located within the partition x.

			lkp, _ := e.lkp.nrg.Load(x)

			// For every energy packet in partition x, add its encoded representation to
			// the new player's fanout buffer.

			for k := range lkp {
				e.buf.Compute(pac.Uid, func(old []byte, _ bool) ([]byte, bool) {
					n, _ := e.mem.nrg.Load(k)
					return append(old, schema.Encode(schema.Food, n.Encode())...), false
				})
			}
		}

		{
			// Search for all the player segments located within the partition x.

			lkp, _ := e.lkp.ply.Load(x)

			// For every player segment in partition x, add its encoded representation to
			// the new player's fanout buffer.

			for k := range lkp {
				e.buf.Compute(pac.Uid, func(old []byte, _ bool) ([]byte, bool) {
					p, _ := e.mem.ply.Load(k)
					return append(old, schema.Encode(schema.Body, p.Vec.Buffer(x))...), false
				})
			}
		}
	}

	// Add the new player to the lookup table based on its currently occupied
	// coordinates.

	for _, x := range ply.Vec.Occupy().Prt {
		e.lkp.ply.Compute(x, func(old map[[2]byte]struct{}, exi bool) (map[[2]byte]struct{}, bool) {
			if !exi {
				old = map[[2]byte]struct{}{}
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
		ply.Vec.Occupy().Prt = nil
	}

	// Add the new player object to the memory table. This ensures that this new
	// player is part of the update loop moving forward.

	{
		e.mem.ply.Store(pac.Uid, ply)
	}

	// After we added the new player to the memory table above, we can calculate
	// the new write deadline that will be enforced in every single fanout cycle.
	// The available duration capacity is divided by the amount of active players
	// that can be processed per active worker process. See engine.New() for the
	// definition of semaphore tickets. See Engine.push() for the respective
	// *time.Timer creation.

	{
		e.tim = timCap(e.mem.ply.Size(), runtime.NumCPU())
	}
}

func timCap(ply int, cpu int) time.Duration {
	if ply <= 1 {
		return Cap / time.Duration(cpu)
	}

	return Cap / time.Duration(cpu) / time.Duration(ply)
}
