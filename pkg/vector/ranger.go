package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

func (v *Vector) Ranger(fnc func(matrix.Coordinate)) {
	for n := v.hea; n != nil; n = n.prv {
		fnc(n.crd)
	}
}
