package filler

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/random"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don <-chan struct{}
	Log logger.Interface
}

type Filler struct {
	ang *random.Random
	crd *random.Random
	qdr *random.Random
}

func New(c Config) *Filler {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}

	var ang *random.Random
	{
		ang = random.New(random.Config{
			Buf: 500,
			Don: c.Don,
			Log: c.Log,
			Max: 255,
			Min: 0,
		})
	}

	var crd *random.Random
	{
		crd = random.New(random.Config{
			Buf: 3000,
			Don: c.Don,
			Log: c.Log,
			Max: matrix.Max,
			Min: matrix.Min,
		})
	}

	var qdr *random.Random
	{
		qdr = random.New(random.Config{
			Buf: 500,
			Don: c.Don,
			Log: c.Log,
			Max: 4,
			Min: 1,
		})
	}

	return &Filler{
		ang: ang,
		crd: crd,
		qdr: qdr,
	}
}

func (f *Filler) Bucket() matrix.Bucket {
	return matrix.Bucket{
		f.crd.Random(), // x0
		f.crd.Random(), // y0
		f.crd.Random(), // x1
		f.crd.Random(), // y1
	}
}

func (f *Filler) Pixel() matrix.Pixel {
	return matrix.Pixel{
		f.crd.Random(), // x2
		f.crd.Random(), // y2
	}
}

func (f *Filler) Space() matrix.Space {
	return matrix.Space{
		f.qdr.Random(), // quadrant
		f.ang.Random(), // angle
	}
}
