package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/puzpuzpuz/xsync/v3"
)

type lookup struct {
	// nrg key: partition, value: location
	nrg *xsync.MapOf[matrix.Partition, map[matrix.Coordinate]struct{}]
	// ply key: partition, value: player ID
	// TODO:infra we should iterate over the existing Vectors instead of creating another representation of location.
	ply *xsync.MapOf[matrix.Partition, map[byte]struct{}]
}
