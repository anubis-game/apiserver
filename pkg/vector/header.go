package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

func (v *Vector) Header() matrix.Coordinate {
	return v.hea.crd
}
