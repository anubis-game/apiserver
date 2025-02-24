package engine

import (
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/vector"
)

// TODO:infra use a separate race slice indexed by player byte IDs. This works
// because there is only a single sequential writer for this data.

func (e *Engine) race(pac router.Packet) {
	e.mem.ply.Compute(pac.Uid, func(ply *player.Player, _ bool) (*player.Player, bool) {
		var mot vector.Motion
		{
			mot = ply.Vec.Motion().Get()
		}

		// The race command triggers a simple switch. There is no race payload. All we
		// do upon receiving the race signal is to flip a player's velocity between
		// normal and racing.

		if mot.Vlc == vector.Nrm {
			mot.Vlc = vector.Rcn
		} else {
			mot.Vlc = vector.Nrm
		}

		{
			ply.Vec.Motion().Set(mot)
		}

		return ply, false
	})
}
