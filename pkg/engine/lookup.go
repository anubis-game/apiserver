package engine

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/google/uuid"
	"github.com/puzpuzpuz/xsync/v3"
)

// lookup
type lookup struct {
	nrg *xsync.MapOf[object.Object, uuid.UUID]
	ply *xsync.MapOf[object.Object, uuid.UUID]
}
