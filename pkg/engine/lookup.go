package engine

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/google/uuid"
	"github.com/puzpuzpuz/xsync/v3"
)

// lookup
type lookup struct {
	nrg *xsync.MapOf[matrix.Bucket, uuid.UUID]
	ply *xsync.MapOf[matrix.Bucket, uuid.UUID]
}
