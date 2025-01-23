package energy

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

type Energy struct {
	Bck matrix.Bucket
	Pxl matrix.Pixel
	Siz byte
}

func (e Energy) Bytes() []byte {
	var buf [7]byte

	copy(buf[0:4], e.Bck[:])
	copy(buf[4:6], e.Pxl[:])

	buf[6] = e.Siz

	return buf[:]
}

func FromBytes(byt []byte) Energy {
	if len(byt) != 7 {
		panic(fmt.Sprintf("expected 7 energy bytes, got %d", len(byt)))
	}

	var e Energy

	copy(e.Bck[:], byt[0:4])
	copy(e.Pxl[:], byt[4:6])

	e.Siz = byt[6]

	return e
}
