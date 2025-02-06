package energy

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/google/uuid"
)

type Energy struct {
	Obj object.Object
	Siz byte
	Typ byte
	Uid uuid.UUID
}

func (e Energy) Bytes() []byte {
	var buf [8]byte

	byt := e.Obj.Byt()
	copy(buf[0:6], byt[:])

	buf[6] = e.Siz
	buf[7] = e.Typ

	return buf[:]
}

func FromBytes(byt []byte) Energy {
	if len(byt) != 8 {
		panic(fmt.Sprintf("expected 8 energy bytes, got %d", len(byt)))
	}

	return Energy{
		Obj: object.FromBytes(byt[0:6]),
		Siz: byt[6],
		Typ: byt[7],
	}
}
