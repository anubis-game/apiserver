package player

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/vector"
)

type Config struct {
	Cli *client.Client
	Uid byte
	Vec *vector.Vector
}

type Player struct {
	Cli *client.Client
	Vec *vector.Vector

	// uid contains this player's wallet specific fanout buffer, containing the
	// action, UID and wallet bytes. This buffer is cached during player
	// initialization so that we do not have to do the same computation all over
	// again every time.
	uid []byte
}

func New(c Config) *Player {
	var uid []byte
	{
		uid = make([]byte, 22)
	}

	{
		uid[0] = byte(schema.Uuid)
		uid[1] = c.Uid
	}

	{
		copy(uid[2:], c.Cli.Wallet().Bytes())
	}

	return &Player{
		Cli: c.Cli,
		Vec: c.Vec,

		uid: uid,
	}
}
