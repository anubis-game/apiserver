package player

import "github.com/anubis-game/apiserver/pkg/matrix"

type Player struct {
	Bck matrix.Bucket
	Pxl matrix.Pixel
	Siz byte
}

func (p Player) Bytes() []byte {
	return []byte{
		p.Bck[0], // x0
		p.Bck[0], // y0
		p.Bck[0], // x1
		p.Bck[0], // y1

		p.Pxl[0], // x2
		p.Pxl[1], // y2

		p.Siz,
	}
}
