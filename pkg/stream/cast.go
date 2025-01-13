package stream

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Stream) cast(_ common.Address, _ *client.Client, inp []byte) error {
	s.cli.Ranger(func(_ common.Address, val *client.Client) {
		val.Stream(inp)
	})

	return nil
}
