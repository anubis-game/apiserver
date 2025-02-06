package object

import "github.com/anubis-game/apiserver/pkg/matrix"

const (
	X0 = 0
	Y0 = 1
	X1 = 2
	Y1 = 3
	X2 = 4
	Y2 = 5
)

// Object contains the coordinates X and Y, describing a precise pixel location
// within the coordinate system.
type Object struct {
	X int
	Y int
}

func FromBytes(byt []byte) Object {
	return Object{
		X: (int(byt[X0]) * 4096) + (int(byt[X1]) * 64) + int(byt[X2]),
		Y: (int(byt[Y0]) * 4096) + (int(byt[Y1]) * 64) + int(byt[Y2]),
	}
}

// Byt returns a precise representation of Object compressed to 6 bytes. The
// encoded version may be used for communicating coordinates over network
// transport.
func (o Object) Byt() [6]byte {
	return [6]byte{
		byte(o.X / 4096),
		byte(o.Y / 4096),
		byte((o.X % 4096) / 64),
		byte((o.Y % 4096) / 64),
		byte(o.X % 64),
		byte(o.Y % 64),
	}
}

// Prt returns the partitioned representation of this Object, which helps to
// group coordinates for caching purposes.
func (o Object) Prt() Object {
	return Object{X: o.X / matrix.Prt * matrix.Prt, Y: o.Y / matrix.Prt * matrix.Prt}
}
