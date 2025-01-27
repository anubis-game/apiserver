package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/google/uuid"
)

// TODO move must also adapt the client window coordinates
func (h *Handler) move(_ uuid.UUID, _ *client.Client, _ []byte) error {
	return nil
}
