package engine

import (
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/vector"
)

func (e *Engine) race(pac router.Packet) {
	var ply *player.Player
	{
		ply, _ = e.mem.ply.Load(pac.Uid)
	}

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
}
