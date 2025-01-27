package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) join(cli *client.Client, _ []byte) error {
	var err error

	{
		err = h.rtr.Create(cli)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
