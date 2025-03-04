package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

func (v *Vector) ranger(fnc func(matrix.Coordinate)) {
	cur := v.hea
	for cur != nil {
		fnc(cur.crd)
		cur = cur.prv
	}
}
