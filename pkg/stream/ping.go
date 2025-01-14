package stream

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/ethereum/go-ethereum/common"
)

var (
	pong = []byte{byte(schema.Pong)}
)

func (s *Stream) ping(_ common.Address, cli *client.Client, _ []byte) error {
	{
		cli.Stream(pong)
	}

	return nil
}
