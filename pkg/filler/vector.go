package filler

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/vector"
)

func (f *Filler) Vector() *vector.Vector {
	return <-f.vec
}

// vector is to simply prepare randomized Vector instances in advance.
func (f *Filler) vector() *vector.Vector {
	// Create a new Motion object so we can point new players towards a randomized
	// direction.

	var mot vector.Motion
	{
		mot = vector.Motion{
			Qdr: byte(f.qdr.Random()),
			Agl: byte(f.agl.Random()),
		}
	}

	// Create a new Vector instance using a random head node.

	var vec *vector.Vector
	{
		vec = vector.New(vector.Config{
			Hea: matrix.Coordinate{
				X: f.crd.Random(),
				Y: f.crd.Random(),
			},
			Mot: mot,
		})
	}

	// We initialize the head of the new vector above with a single coordinate
	// object. Below we use this head segment as basis for the Vector's expansion.
	// 1 head plus 4 expansions gives us a vector with 5 segments lined up towards
	// the same direction, because we use the same motion configuration every
	// time.

	for range 4 {
		vec.Update(int(vector.Si/vector.Li), mot.Qdr, mot.Agl, vector.Nrm)
	}

	return vec
}
