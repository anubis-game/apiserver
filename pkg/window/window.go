package window

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

const (
	//
	Siz byte = 21
)

type Config struct {
	Bck matrix.Bucket
	Pxl matrix.Pixel
	Spc matrix.Space
}

type Window struct {
	win [2]matrix.Bucket
	pxl matrix.Pixel
	spc matrix.Space
}

func New(c Config) *Window {
	return &Window{
		win: [2]matrix.Bucket{
			c.Bck,
			c.Bck.Scale(Siz),
		},
		pxl: c.Pxl,
		spc: c.Spc,
	}
}

func (w *Window) Bytes() []byte {
	return []byte{
		// bottom left
		w.win[0][0], // x0
		w.win[0][1], // y0
		w.win[0][2], // x1
		w.win[0][3], // y1

		// top right
		w.win[1][0], // x0
		w.win[1][1], // y0
		w.win[1][2], // x1
		w.win[1][3], // y1

		// pixel location
		w.pxl[0], // x2
		w.pxl[1], // y2

		// direction
		w.spc[0], // quadrant
		w.spc[1], // angle
	}
}

func (w *Window) Pixel() matrix.Pixel {
	return w.pxl
}

func (w *Window) Space() matrix.Space {
	return w.spc
}

func (w *Window) Window() [2]matrix.Bucket {
	return w.win
}
