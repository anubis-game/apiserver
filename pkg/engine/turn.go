package engine

import (
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// TODO:infra use a separate Direct slice indexed by player byte IDs. This works
// because there is only a single sequential writer for this data.

// TODO:infra motion and charax structs should not even be part of the
// player/vector objects.

func (e *Engine) turn(pac router.Packet) {
	// If we do not receive exactly two bytes, then we simply ignore the user
	// input. The two required bytes here are the quadrant byte and the angle
	// byte.

	if len(pac.Byt) != 2 {
		return
	}

	// If the quadrant byte is not one of [1 2 3 4], then we simply ignore the
	// user input.

	if pac.Byt[0]-1 > 3 {
		return
	}

	e.mem.ply.Compute(pac.Uid, func(ply *player.Player, _ bool) (*player.Player, bool) {
		var mot vector.Motion
		{
			mot = ply.Vec.Motion().Get()
		}

		{
			mot.Qdr = pac.Byt[0]
			mot.Agl = pac.Byt[1]
		}

		{
			ply.Vec.Motion().Set(mot)
		}

		return ply, false
	})
}
