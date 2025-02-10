package engine

import (
	"github.com/google/uuid"
	"github.com/puzpuzpuz/xsync/v3"
)

// buffer contains various messages prepared to be sent out to all connected
// clients during the time based fanout procedure. The UUID key in the maps
// below defines the connected clients to which the associated message bytes
// should be sent to. The buffers here can only be used for new players joining
// the game, because only for those situations can we prepare a full buffer
// representation ahead of time.
type buffer struct {
	nrg *xsync.MapOf[uuid.UUID, [][]byte]
	ply *xsync.MapOf[uuid.UUID, [][]byte]
}
