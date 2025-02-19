package energy

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/schema"
)

type Config struct {
	Obj object.Object
	Siz byte
	Typ byte
}

type Energy struct {
	Obj object.Object
	Siz byte
	Typ byte

	nrg []byte
}

func New(c Config) *Energy {
	var nrg []byte
	{
		nrg = make([]byte, 9)
	}

	{
		nrg[0] = byte(schema.Food)
	}

	{
		b := c.Obj.Byt()
		copy(nrg[1:1+object.Len], b[:])
	}

	{
		nrg[7] = c.Siz
		nrg[8] = c.Typ
	}

	return &Energy{
		Obj: c.Obj,
		Siz: c.Siz,
		Typ: c.Typ,

		nrg: nrg,
	}
}
