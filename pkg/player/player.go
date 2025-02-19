package player

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/anubis-game/apiserver/pkg/vector"
)

type Config struct {
	Cli *client.Client
	Uid [2]byte
	Vec *vector.Vector
}

type Player struct {
	Cli *client.Client
	Vec *vector.Vector

	buf setter.Interface[[]byte]
	wal []byte
}

func New(c Config) *Player {
	var wal []byte
	{
		wal = make([]byte, 22)

		copy(wal[:2], c.Uid[:])
		copy(wal[2:], c.Cli.Wallet().Bytes())
	}

	return &Player{
		Cli: c.Cli,
		Vec: c.Vec,

		buf: setter.New[[]byte](),
		wal: wal,
	}
}
