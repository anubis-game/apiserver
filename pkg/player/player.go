package player

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/anubis-game/apiserver/pkg/window"
	"github.com/google/uuid"
)

const (
	// Rad is the initial radius of a player's head and body parts.
	Rad byte = 10
	// Siz is the initial amount of points that a player is worth.
	Siz byte = 50
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
	Win *window.Window
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
		Win: window.New(window.Config{
			Obj: c.Vec.Header(),
		}),
	}
}

func (p Player) Bytes() []byte {
	vec := p.Vec.Bytes()
	byt := make([]byte, 22+len(vec))
	mot := p.Vec.Motion().Get()

	copy(byt[0:16], p.Uid[:])

	byt[16] = p.Crx.Rad
	byt[17] = p.Crx.Siz
	byt[18] = p.Crx.Typ

	byt[19] = mot.AGL
	byt[20] = mot.QDR
	byt[21] = mot.VLC

	copy(byt[22:], vec[:])

	return byt[:]
}

func FromBytes(byt []byte) Player {
	if len(byt) < 22 {
		panic(fmt.Sprintf("expected at least 22 player bytes, got %d", len(byt)))
	}

	var p Player

	copy(p.Uid[:], byt[0:16])

	p.Crx.Rad = byt[16]
	p.Crx.Siz = byt[17]
	p.Crx.Typ = byt[18]

	p.Vec = vector.FromBytes(byt[22:])

	p.Vec.Motion().Set(vector.Motion{
		AGL: byt[19],
		QDR: byt[20],
		VLC: byt[21],
	})

	return p
}
