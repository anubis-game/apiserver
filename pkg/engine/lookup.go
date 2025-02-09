package engine

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/google/uuid"
	"github.com/puzpuzpuz/xsync/v3"
)

// TODO lookup needs to key partition coordinates
type lookup struct {
	nrg *xsync.MapOf[object.Object, map[uuid.UUID]struct{}]
	ply *xsync.MapOf[object.Object, map[uuid.UUID]struct{}]
}
