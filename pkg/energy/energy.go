package energy

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

const (
	// EnergyBytes is used to derive the byte slice length of a single energy
	// packet.
	EnergyBytes = 8
)

type Config struct {
	Crd matrix.Coordinate
	Siz byte
	Typ byte
}

func New(c Config) []byte {
	var byt []byte
	{
		byt = make([]byte, EnergyBytes)
	}

	{
		b := c.Crd.Byt()
		copy(byt[:matrix.CoordinateBytes], b[:])
	}

	{
		b := [2]byte{c.Siz, c.Typ}
		copy(byt[matrix.CoordinateBytes:], b[:])
	}

	return byt
}
