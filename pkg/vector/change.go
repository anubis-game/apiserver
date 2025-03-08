package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

func (v *Vector) Change(prt matrix.Partition) []matrix.Coordinate {
	return v.ocd[prt]
}
