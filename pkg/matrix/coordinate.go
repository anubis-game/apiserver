package matrix

const (
	X0 = 0
	Y0 = 1
	X1 = 2
	Y1 = 3
	X2 = 4
	Y2 = 5
)

const (
	CoordinateBytes = 6
)

// Coordinate contains the coordinates X and Y, describing a precise pixel
// location within the coordinate system.
type Coordinate struct {
	X int
	Y int
}

// TODO:infra instead of computing the coordinate partitions and byte
// representations over and over again, we should generate them once.

func NewCoordinate(byt []byte) Coordinate {
	return Coordinate{
		X: (int(byt[X0]) * 4096) + (int(byt[X1]) * 64) + int(byt[X2]),
		Y: (int(byt[Y0]) * 4096) + (int(byt[Y1]) * 64) + int(byt[Y2]),
	}
}

// Byt returns a precise representation of Coordinate compressed to 6 bytes.
// Converting two integer coordinates into buckets of 6 bytes is about 10 times
// faster than doing big endian byte shifting. This is because byte shifting
// requires 1 allocation per integer, and every integer occupies 4 bytes. Our
// approach below does not allocate and saves 2 bytes, which is also relevant
// for communicating coordinates over network transport.
func (c Coordinate) Byt() [CoordinateBytes]byte {
	return [CoordinateBytes]byte{
		byte(c.X / 4096),
		byte(c.Y / 4096),
		byte((c.X % 4096) / 64),
		byte((c.Y % 4096) / 64),
		byte(c.X % 64),
		byte(c.Y % 64),
	}
}

// Prt returns the partitioned representation of this Coordinate, which helps to
// group coordinates for caching purposes.
func (c Coordinate) Prt() Partition {
	return Partition{
		X: (c.X / Prt) * Prt,
		Y: (c.Y / Prt) * Prt,
	}
}
