package object

const (
	X0 = 0
	Y0 = 1
	X1 = 2
	Y1 = 3
	X2 = 4
	Y2 = 5
)

type Bucket [6]byte

func (b Bucket) Object() Object {
	return Object{
		X: (int(b[X0]) * 4096) + (int(b[X1]) * 64) + int(b[X2]),
		Y: (int(b[Y0]) * 4096) + (int(b[Y1]) * 64) + int(b[Y2]),
	}
}
