package engine

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/puzpuzpuz/xsync/v3"
)

type lookup struct {
	// nrg key: partition, value: location
	nrg *xsync.MapOf[object.Object, map[object.Object]struct{}]
	// ply key: partition, value: player ID
	ply *xsync.MapOf[object.Object, map[byte]struct{}]
}
