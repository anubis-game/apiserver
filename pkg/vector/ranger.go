package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

func (v *Vector) Ranger(f func(matrix.Coordinate)) {
	for n := v.hea; n != nil; n = n.prv {
		f(n.crd)
	}
}
