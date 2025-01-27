package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/google/uuid"
	"github.com/xh3b4sd/tracer"
)

func (h *Handler) join(uid uuid.UUID, cli *client.Client, _ []byte) error {
	var err error

	{
		err = h.rtr.Create(uid, cli)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
