package energy

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/schema"
)

type Config struct {
	Crd matrix.Coordinate
	Siz byte
	Typ byte
}

func New(c Config) []byte {
	var nrg []byte
	{
		nrg = make([]byte, 9)
	}

	{
		nrg[0] = byte(schema.Food)
	}

	{
		b := c.Crd.Byt()
		copy(nrg[1:1+matrix.CoordinateBytes], b[:])
	}

	{
		nrg[7] = c.Siz
		nrg[8] = c.Typ
	}

	return nrg
}
