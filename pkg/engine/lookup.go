package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

type lookup struct {
	nrg map[matrix.Partition]map[matrix.Coordinate]struct{}
	pt1 map[matrix.Partition]map[byte]struct{}
	pt8 map[matrix.Partition]map[byte]struct{}
}
