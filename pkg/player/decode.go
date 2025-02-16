package player

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/vector"
)

func Decode(byt []byte) Player {
	if len(byt) < 8 {
		panic(fmt.Sprintf("expected at least 8 player bytes, got %d", len(byt)))
	}

	var p Player

	copy(p.Uid[:], byt[:2])

	p.Vec = vector.Decode(byt[8:])

	p.Vec.Charax().Set(vector.Charax{
		Rad: int(byt[2]),
		Siz: int(byt[3]),
		Typ: byt[4],
	})

	p.Vec.Motion().Set(vector.Motion{
		Qdr: byt[5],
		Agl: byt[6],
		Vlc: byt[7],
	})

	return p
}
