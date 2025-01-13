package stream

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Stream) join(wal common.Address, cli *client.Client, _ []byte) error {
	// At first we want to check whether the given Wallet address is already part
	// of the braodcasting worker pool, because we do not want to do unnecessary
	// work.

	{
		exi := s.cli.Exists(wal)
		if exi {
			s.log.Log(
				s.ctx,
				"level", "warning",
				"message", "already joined",
				"wallet", wal.String(),
			)

			return nil
		}
	}

	// Upon joining, we add the user to the broadcasting system in order to
	// provide them with realtime data primitives over the given client
	// connection.

	{
		s.cli.Update(wal, cli)
	}

	var out []byte
	{
		out = schema.Encode(schema.Join, wal.Bytes())
	}

	s.cli.Ranger(func(_ common.Address, val *client.Client) {
		val.Stream(out)
	})

	return nil
}
