package engine

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/google/uuid"
)

type memory struct {
	// nrg
	nrg map[uuid.UUID]*energy.Energy
	// ply contains all active player information, including their connected
	// clients. The native Go map is synchronized via channel access in line with
	// the time based fanout procedure for minimum latency.
	ply map[uuid.UUID]*player.Player
}
