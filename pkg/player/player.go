package player

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/google/uuid"
)

const (
	// Siz describes half of the initial window size along x and y axis. The goal
	// is to put the player into the middle of this window, which means that we
	// have to define the edges and the center of the window. E.g. a size of 5
	// implies the total window length along x and y axis of 11 inner buckets,
	// which puts the player into the middle of the window at the relative
	// coordinates x=5 y=5. The player has then 5 inner buckets all around the
	// inner bucket that the player is put into.
	Siz byte = 5
)

type Config struct {
	Bck matrix.Bucket
	Cli *client.Client
	Pxl matrix.Pixel
	Spc matrix.Space
	Uid uuid.UUID
}

type Player struct {
	Cli *client.Client
	Obj matrix.Object
	Spc matrix.Space
	Win matrix.Window
}

func New(c Config) *Player {
	return &Player{
		Cli: c.Cli,
		Obj: matrix.Object{
			Bck: c.Bck,
			Pxl: c.Pxl,
			Pro: matrix.Profile{
				Siz, // size
				0,   // type
			},
			Uid: c.Uid,
		},
		Spc: c.Spc,
		Win: matrix.Window{
			c.Bck.Dec(Siz), // bottom left
			c.Bck.Inc(Siz), // top right
		},
	}
}

func (p Player) Bytes() []byte {
	var buf [26]byte

	copy(buf[0:4], p.Obj.Bck[:])
	copy(buf[4:6], p.Obj.Pxl[:])
	copy(buf[6:8], p.Obj.Pro[:])
	copy(buf[8:24], p.Obj.Uid[:])
	copy(buf[24:26], p.Spc[:])

	return buf[:]
}

func FromBytes(byt []byte) Player {
	if len(byt) != 26 {
		panic(fmt.Sprintf("expected 26 player bytes, got %d", len(byt)))
	}

	var p Player

	copy(p.Obj.Bck[:], byt[0:4])
	copy(p.Obj.Pxl[:], byt[4:6])
	copy(p.Obj.Pro[:], byt[6:8])
	copy(p.Obj.Uid[:], byt[8:24])
	copy(p.Spc[:], byt[24:26])

	return p
}
