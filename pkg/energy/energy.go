package energy

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

type Energy struct {
	Obj matrix.Object
}

func (e Energy) Bytes() []byte {
	var buf [8]byte

	copy(buf[0:4], e.Obj.Bck[:])
	copy(buf[4:6], e.Obj.Pxl[:])
	copy(buf[6:8], e.Obj.Pro[:])

	return buf[:]
}

func FromBytes(byt []byte) Energy {
	if len(byt) != 8 {
		panic(fmt.Sprintf("expected 8 energy bytes, got %d", len(byt)))
	}

	var e Energy

	copy(e.Obj.Bck[:], byt[0:4])
	copy(e.Obj.Pxl[:], byt[4:6])
	copy(e.Obj.Pro[:], byt[6:8])

	return e
}
