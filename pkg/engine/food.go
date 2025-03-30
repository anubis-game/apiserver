package engine

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/matrix"
)

func (e *Engine) food() {
	var s byte

	if e.cyc%10 == 0 {
		s = 2 // every 250 ms
	} else if e.cyc%20 == 0 {
		s = 5 // every 500 ms
	} else if e.cyc%40 == 0 {
		s = 10 // every 1000 ms
	} else {
		return // don't generate a new energy packet
	}

	var c matrix.Coordinate
	{
		c = matrix.Coordinate{
			X: e.crd.Random(),
			Y: e.crd.Random(),
		}
	}

	var n []byte
	{
		n = energy.New(energy.Config{
			Crd: c,
			Siz: s,
			Typ: 0, // TODO:infra randomize energy type
		})
	}

	{
		e.mem.nrg[c] = n
	}

	// TODO find all connected players that can see the new energy packet
}
