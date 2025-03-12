package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

type Change struct {
	Hea matrix.Coordinate
	Tai matrix.Coordinate // TODO:infra what do we even need this new tail change for?
	Rem []matrix.Coordinate
}

func (v *Vector) Change() Change {
	return v.ocd
}
