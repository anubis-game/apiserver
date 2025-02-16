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
	agl *random.Random
	crd *random.Random
	don <-chan struct{}
	qdr *random.Random
	vec chan Vector
}

func New(c Config) *Filler {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}

	var agl *random.Random
	{
		agl = random.New(random.Config{
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
			Buf: 1000,
			Don: c.Don,
			Log: c.Log,
			Max: matrix.Max - matrix.Thr,
			Min: matrix.Min + matrix.Thr,
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
		agl: agl,
		crd: crd,
		qdr: qdr,
		vec: make(chan Vector, 500),
	}
}
