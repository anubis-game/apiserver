package player

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/google/uuid"
)

const (
	// Rad is the initial radius of a player's head and body parts.
	Rad = 10
	// Siz is the initial amount of points that a player is worth.
	Siz = 50
)

type Config struct {
	Cli *client.Client
	Uid uuid.UUID
	Vec *vector.Vector
}

type Player struct {
	Cli *client.Client
	Crx Charax
	Uid uuid.UUID
	Vec *vector.Vector
}

func New(c Config) *Player {
	return &Player{
		Cli: c.Cli,
		Crx: Charax{
			Rad: Rad,
			Siz: Siz,
			Typ: 0, // TODO randomize or configure the player suit based on the user's preference
		},
		Uid: c.Uid,
		Vec: c.Vec,
	}
}

func (p Player) Bytes() []byte {
	var buf [26]byte

	copy(buf[0:4], p.Bck[:])
	copy(buf[4:6], p.Pxl[:])
	copy(buf[6:8], p.Pro[:])
	copy(buf[8:24], p.Uid[:])
	copy(buf[24:26], p.Spc[:])

	return buf[:]
}

func FromBytes(byt []byte) Player {
	if len(byt) != 26 {
		panic(fmt.Sprintf("expected 26 player bytes, got %d", len(byt)))
	}

	var p Player

	copy(p.Bck[:], byt[0:4])
	copy(p.Pxl[:], byt[4:6])
	copy(p.Pro[:], byt[6:8])
	copy(p.Uid[:], byt[8:24])
	copy(p.Spc[:], byt[24:26])

	return p
}
