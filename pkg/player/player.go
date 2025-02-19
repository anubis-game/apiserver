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

	// buf contains various messages prepared to be sent out to this player's
	// client during the time based fanout procedure. The byte slice may contain
	// multiple encoded messages.
	//
	//     [ 1 action byte ] [ N buffer bytes ] [ 1 action byte ] [ N buffer bytes ] [ 1 action byte ] ...
	//
	buf setter.Interface[[]byte]
	// wal contains this player's wallet specific fanout buffer, containing the
	// UID and wallet bytes. This buffer is cached during player initialization so
	// that we do not have to do the same computation all over again every time.
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
