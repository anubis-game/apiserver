package energy

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/google/uuid"
)

type Energy struct {
	Bck matrix.Bucket
	Pxl matrix.Pixel
	Pro matrix.Profile
	Uid uuid.UUID
}

func (e Energy) Bytes() []byte {
	var buf [8]byte

	copy(buf[0:4], e.Bck[:])
	copy(buf[4:6], e.Pxl[:])
	copy(buf[6:8], e.Pro[:])

	return buf[:]
}

func FromBytes(byt []byte) Energy {
	if len(byt) != 8 {
		panic(fmt.Sprintf("expected 8 energy bytes, got %d", len(byt)))
	}

	var e Energy

	copy(e.Bck[:], byt[0:4])
	copy(e.Pxl[:], byt[4:6])
	copy(e.Pro[:], byt[6:8])

	return e
}
