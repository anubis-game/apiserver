package engine

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/google/uuid"
)

type memory struct {
	// cli contains all connected clients. This is a native Go map, and we
	// synchronize it via channel access in line with the time based fanout
	// procedure.
	cli map[uuid.UUID]*client.Client
	// nrg
	nrg map[uuid.UUID]*energy.Energy
	// ply
	ply map[uuid.UUID]*player.Player
}
