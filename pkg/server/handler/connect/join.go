package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/google/uuid"
)

func (h *Handler) join(uid uuid.UUID, cli *client.Client, _ []byte) error {
	return h.rtr.Join(uid, cli, nil)
}
