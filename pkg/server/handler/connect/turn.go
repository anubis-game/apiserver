package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
)

func (h *Handler) turn(uid byte, cli *client.Client, byt []byte) error {
	return h.rtr.Turn(uid, cli, byt)
}
