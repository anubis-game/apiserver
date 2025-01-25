package connect

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
)

var (
	pong = []byte{byte(schema.Pong)}
)

func (h *Handler) ping(cli *client.Client, _ []byte) error {
	{
		cli.Stream(pong)
	}

	return nil
}
