package player

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/schema"
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

	// uid contains this player's wallet specific fanout buffer, containing the
	// action, UID and wallet bytes. This buffer is cached during player
	// initialization so that we do not have to do the same computation all over
	// again every time.
	uid []byte
}

func New(c Config) *Player {
	var uid []byte
	{
		uid = make([]byte, 23)
	}

	{
		uid[0] = byte(schema.Uuid)
	}

	{
		copy(uid[1:3], c.Uid[:])
		copy(uid[3:], c.Cli.Wallet().Bytes())
	}

	return &Player{
		Cli: c.Cli,
		Vec: c.Vec,

		uid: uid,
	}
}
