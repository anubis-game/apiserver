package player

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/ethereum/go-ethereum/common"
)

type Player struct {
	Bck matrix.Bucket
	Pxl matrix.Pixel
	Siz byte
	Wal common.Address
}

func (p Player) Bytes() []byte {
	var buf [27]byte

	copy(buf[0:20], p.Wal[:])
	copy(buf[20:24], p.Bck[:])
	copy(buf[24:26], p.Pxl[:])

	buf[26] = p.Siz

	return buf[:]
}

func FromBytes(byt []byte) Player {
	if len(byt) != 27 {
		panic(fmt.Sprintf("expected 27 player bytes, got %d", len(byt)))
	}

	var p Player

	copy(p.Wal[:], byt[0:20])
	copy(p.Bck[:], byt[20:24])
	copy(p.Pxl[:], byt[24:26])

	p.Siz = byt[26]

	return p
}
