package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/google/uuid"
)

func (h *Handler) race(uid uuid.UUID, cli *client.Client, _ []byte) error {
	return h.rtr.Race(uid, cli, nil)
}
