package player

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/google/uuid"
)

type Config struct {
	Cli *client.Client
	Uid uuid.UUID
	Vec *vector.Vector
}

type Player struct {
	Cli *client.Client
	Uid uuid.UUID
	Vec *vector.Vector
}

func New(c Config) *Player {
	return &Player{
		Cli: c.Cli,
		Uid: c.Uid,
		Vec: c.Vec,
	}
}

func (p Player) Bytes() []byte {
	vec := p.Vec.Bytes()
	byt := make([]byte, 22+len(vec))
	crx := p.Vec.Charax().Get()
	mot := p.Vec.Motion().Get()

	copy(byt[0:16], p.Uid[:])

	byt[16] = crx.Rad
	byt[17] = crx.Siz
	byt[18] = crx.Typ

	byt[19] = mot.Qdr
	byt[20] = mot.Agl
	byt[21] = mot.Vlc

	copy(byt[22:], vec[:])

	return byt[:]
}

func FromBytes(byt []byte) Player {
	if len(byt) < 22 {
		panic(fmt.Sprintf("expected at least 22 player bytes, got %d", len(byt)))
	}

	var p Player

	copy(p.Uid[:], byt[0:16])

	p.Vec = vector.FromBytes(byt[22:])

	p.Vec.Charax().Set(vector.Charax{
		Rad: byt[16],
		Siz: byt[17],
		Typ: byt[18],
	})

	p.Vec.Motion().Set(vector.Motion{
		Qdr: byt[19],
		Agl: byt[20],
		Vlc: byt[21],
	})

	return p
}
