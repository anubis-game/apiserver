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
	// Create random quadrant and angle bytes, so we can point new players towards
	// a randomized direction.

	var qdr byte
	var agl byte
	{
		qdr = byte(f.qdr.Random())
		agl = byte(f.agl.Random())
	}

	// Create a new Vector instance using a random head node.

	var vec *vector.Vector
	{
		vec = vector.New(vector.Config{
			Hea: matrix.Coordinate{
				X: f.crd.Random(),
				Y: f.crd.Random(),
			},
			Mot: vector.Motion{
				Qdr: qdr,
				Agl: agl,
			},
		})
	}

	// We initialize the head of the new Vector above with a single coordinate
	// object. Below we use this head node as basis for the Vector's expansion. 1
	// head plus 4 expansions gives us a Vector with 5 nodes, all lined up towards
	// the same direction, because we use the same motion configuration every
	// time.

	for range 4 {
		vec.Update(int(vector.Si/vector.Li), qdr, agl, vector.Nrm)
	}

	return vec
}
