package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
)

func (h *Handler) race(uid byte, cli *client.Client, _ []byte) error {
	return h.rtr.Race(uid, cli, nil)
}
