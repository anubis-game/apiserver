package player

import (
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Uid byte
	Vec *vector.Vector
	Wal common.Address
}

// TODO:infra remove the concept of a player. All we need is to track the Vector
// and Wallet.
type Player struct {
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
		copy(uid[2:], c.Wal.Bytes())
	}

	return &Player{
		Vec: c.Vec,

		uid: uid,
	}
}
