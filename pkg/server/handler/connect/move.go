package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
)

func (h *Handler) move(uid [2]byte, cli *client.Client, byt []byte) error {
	return h.rtr.Move(uid, cli, byt)
}
