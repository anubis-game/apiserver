package engine

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/object"
)

func (e *Engine) food(nrg *energy.Energy) {
	{
		e.men.Store(nrg.Obj, nrg)
	}

	e.pen.Compute(nrg.Obj.Prt(), func(old map[object.Object]struct{}, exi bool) (map[object.Object]struct{}, bool) {
		if !exi {
			old = map[object.Object]struct{}{}
		}

		{
			old[nrg.Obj] = struct{}{}
		}

		return old, false
	})
}
