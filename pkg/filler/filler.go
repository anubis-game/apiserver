package filler

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/random"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Cap int
	Don <-chan struct{}
	Log logger.Interface
}

type Filler struct {
	agl *random.Random
	crd *random.Random
	don <-chan struct{}
	qdr *random.Random
	vec chan *vector.Vector
}

func New(c Config) *Filler {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}

	var agl *random.Random
	{
		agl = random.New(random.Config{
			Buf: c.Cap,
			Don: c.Don,
			Log: c.Log,
			Max: 255,
			Min: 0,
		})
	}

	// We generate random coordinates based on a threshold around the edges of the
	// game map, in which players cannot be placed initially, upon joining the
	// game. The purpose of this buffer region is to not put players too close to
	// the edges of the game, so that they cannot run into the wall accidentally.

	var crd *random.Random
	{
		crd = random.New(random.Config{
			Buf: c.Cap * 2,
			Don: c.Don,
			Log: c.Log,
			Max: matrix.Max - matrix.Pt8,
			Min: matrix.Min + matrix.Pt8,
		})
	}

	var qdr *random.Random
	{
		qdr = random.New(random.Config{
			Buf: c.Cap,
			Don: c.Don,
			Log: c.Log,
			Max: 4,
			Min: 1,
		})
	}

	return &Filler{
		agl: agl,
		crd: crd,
		qdr: qdr,
		vec: make(chan *vector.Vector, c.Cap),
	}
}
