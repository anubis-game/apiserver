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

	b := e.Obj.Bucket()
	copy(buf[0:6], b[:])

	buf[6] = e.Siz
	buf[7] = e.Typ

	return buf[:]
}

func FromBytes(byt []byte) Energy {
	if len(byt) != 8 {
		panic(fmt.Sprintf("expected 8 energy bytes, got %d", len(byt)))
	}

	return Energy{
		Obj: object.Bucket(byt[0:6]).Object(),
		Siz: byt[6],
		Typ: byt[7],
	}
}
