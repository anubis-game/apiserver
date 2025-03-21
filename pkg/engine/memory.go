package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

type memory struct {
	// nrg contains all active energy packets currently placed within the game
	// map. Energy is identified by its precise X and Y coordinates, because only
	// one energy packet can be in the same place at the same time. We can refer
	// to energy packets using their position only because energy packets don't
	// move.
	nrg map[matrix.Coordinate][]byte
	// vec contains all player Vectors. Active players have non-nil Vectors.
	vec []*vector.Vector
}

func newMemory(c int) *memory {
	return &memory{
		nrg: make(map[matrix.Coordinate][]byte),
		vec: make([]*vector.Vector, c),
	}
}
