package window

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
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
	Pxl matrix.Pixel
	Spc matrix.Space
}

type Window struct {
	wbl matrix.Bucket
	wcn matrix.Bucket
	wtr matrix.Bucket
	pxl matrix.Pixel
	spc matrix.Space
}

func New(c Config) *Window {
	return &Window{
		wbl: c.Bck,
		wcn: c.Bck.Scale(Siz),
		wtr: c.Bck.Scale(Siz * 2),
		pxl: c.Pxl,
		spc: c.Spc,
	}
}

func (w *Window) Bytes() []byte {
	var buf [16]byte

	copy(buf[0:4], w.wbl[:])
	copy(buf[4:8], w.wcn[:])
	copy(buf[8:12], w.wtr[:])
	copy(buf[12:14], w.pxl[:])
	copy(buf[14:16], w.spc[:])

	return buf[:]
}

func (w *Window) Pixel() matrix.Pixel {
	return w.pxl
}

func (w *Window) Space() matrix.Space {
	return w.spc
}

func (w *Window) Window() (matrix.Bucket, matrix.Bucket, matrix.Bucket) {
	return w.wbl, w.wcn, w.wtr
}

func FromBytes(byt []byte) *Window {
	if len(byt) != 16 {
		panic(fmt.Sprintf("expected 16 energy bytes, got %d", len(byt)))
	}

	var w Window

	copy(w.wbl[:], byt[0:4])
	copy(w.wcn[:], byt[4:8])
	copy(w.wtr[:], byt[8:12])
	copy(w.pxl[:], byt[12:14])
	copy(w.spc[:], byt[14:16])

	return &w
}
