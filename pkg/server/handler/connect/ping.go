package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/google/uuid"
)

var (
	pong = []byte{byte(schema.Pong)}
)

func (h *Handler) ping(_ uuid.UUID, cli *client.Client, _ []byte) error {
	{
		cli.Stream(pong)
	}

	return nil
}
