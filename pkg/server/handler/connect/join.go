package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
)

func (h *Handler) join(uid [2]byte, cli *client.Client, _ []byte) error {
	return h.rtr.Join(uid, cli, nil)
}
