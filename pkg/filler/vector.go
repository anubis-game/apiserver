package filler

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/vector"
)

func (f *Filler) Vector() *vector.Vector {
	return <-f.vec
}

func (f *Filler) vector() *vector.Vector {
	// vector.New() allocates a full slice of X and Y coordinates. Allocating most
	// of the vectors required at runtime helps us to increase runtime
	// performance, because the vector allocation is not done within performance
	// sensitive contexts.

	var vec *vector.Vector
	{
		vec = vector.New(vector.Config{
			Mot: vector.Motion{
				Qdr: byte(f.qdr.Random()),
				Agl: byte(f.ang.Random()),
				Vlc: 0x1,
			},
			Obj: []object.Object{
				{
					X: f.crd.Random(),
					Y: f.crd.Random(),
				},
			},
		})
	}

	// We initialize the head of the new vector above with a single object.  Below
	// we use this head as basis for vector expansion. 1 head plus 4 expansions
	// gives us a vector with 5 segments lined up towards the same direction,
	// because we use the same motion configuration every time.

	{
		vec.Expand(vec.Target(vec.Motion().Get()))
		vec.Expand(vec.Target(vec.Motion().Get()))
		vec.Expand(vec.Target(vec.Motion().Get()))
		vec.Expand(vec.Target(vec.Motion().Get()))
	}

	return vec
}
