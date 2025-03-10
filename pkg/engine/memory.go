package engine

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/puzpuzpuz/xsync/v3"
)

type memory struct {
	// nrg contains all active energy packets currently placed within the game
	// map. Energy is identified by its precise X and Y coordinates, because only
	// one energy packet can be in the same place at the same time. We can refer
	// to energy packets using their position only because energy packets don't
	// move.
	nrg *xsync.MapOf[matrix.Coordinate, *energy.Energy]
	// vec contains all Vectors for the mapped byte ID.
	vec *xsync.MapOf[byte, *vector.Vector]
}
