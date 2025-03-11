package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

type Change struct {
	Hea matrix.Coordinate
	Tai matrix.Coordinate
	Rem []matrix.Coordinate
}

func (v *Vector) Change() Change {
	return v.ocd
}
