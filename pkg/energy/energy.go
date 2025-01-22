package energy

import "github.com/anubis-game/apiserver/pkg/matrix"

type Energy struct {
	Bck matrix.Bucket
	Pxl matrix.Pixel
	Siz byte
}

func (e Energy) Bytes() []byte {
	return []byte{
		e.Bck[0], // x0
		e.Bck[0], // y0
		e.Bck[0], // x1
		e.Bck[0], // y1

		e.Pxl[0], // x2
		e.Pxl[1], // y2

		e.Siz,
	}
}
