package engine

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/player"
)

type memory struct {
	// nrg contains all active energy packets currently placed within the game
	// map. Energy is identified by its precise X and Y coordinates, because only
	// one energy packet can be in the same place at the same time. We can refer
	// to energy packets using their position only, because energy packets don't
	// move.
	nrg map[object.Object]*energy.Energy
	// ply contains all active player information, including their connected
	// clients. The native Go map is synchronized via channel access in line with
	// the time based fanout procedure for minimum latency.
	ply map[[2]byte]*player.Player
}
