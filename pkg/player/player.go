package player

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/vector"
)

type Config struct {
	Cli *client.Client
	Uid [2]byte
	Vec *vector.Vector
}

type Player struct {
	Cli *client.Client
	Uid [2]byte
	Vec *vector.Vector

	ply []byte
	wal []byte
}

func New(c Config) *Player {
	vec := c.Vec.Encode()
	crx := c.Vec.Charax().Get()
	mot := c.Vec.Motion().Get()

	var ply []byte
	{
		ply = make([]byte, 8+len(vec))

		copy(ply[0:2], c.Uid[:])
		copy(ply[8:], vec)

		ply[2] = byte(crx.Rad)
		ply[3] = byte(crx.Siz)
		ply[4] = crx.Typ

		ply[5] = mot.Qdr
		ply[6] = mot.Agl
		ply[7] = mot.Vlc
	}

	var wal []byte
	{
		wal = make([]byte, 22)

		copy(wal[:2], c.Uid[:])
		copy(wal[2:], c.Cli.Wallet().Bytes())
	}

	return &Player{
		Cli: c.Cli,
		Uid: c.Uid,
		Vec: c.Vec,

		ply: ply,
		wal: wal,
	}
}
