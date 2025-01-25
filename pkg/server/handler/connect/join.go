package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
)

func (h *Handler) join(cli *client.Client, _ []byte) error {
	{
		h.rtr.Create(cli)
	}

	return nil
}
