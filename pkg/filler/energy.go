package filler

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/object"
)

type Energy struct {
	Obx int
	Oby int
	Typ int
}

func (f *Filler) Energy(siz byte) *energy.Energy {
	var nfl Energy
	{
		nfl = <-f.nrg
	}

	var nrg *energy.Energy
	{
		nrg = energy.New(energy.Config{
			Obj: object.Object{
				X: nfl.Obx,
				Y: nfl.Oby,
			},
			Siz: siz,
			Typ: byte(nfl.Typ),
		})
	}

	return nrg
}

// energy is to simply prepare randomized Energy configuration in advance.
func (f *Filler) energy() Energy {
	return Energy{
		Obx: f.crd.Random(),
		Oby: f.crd.Random(),
		Typ: f.qdr.Random(),
	}
}
