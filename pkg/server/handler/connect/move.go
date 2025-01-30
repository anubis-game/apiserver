package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/google/uuid"
)

func (h *Handler) move(uid uuid.UUID, cli *client.Client, byt []byte) error {
	return h.rtr.Move(uid, cli, byt)
}
