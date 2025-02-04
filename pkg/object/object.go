package object

// Object contains the coordinates X and Y, describing a precise pixel location
// within the coordinate system.
type Object struct {
	X int
	Y int
}

func (o Object) Bucket() Bucket {
	return Bucket{
		byte(o.X / 4096),
		byte(o.Y / 4096),
		byte((o.X % 4096) / 64),
		byte((o.Y % 4096) / 64),
		byte(o.X % 64),
		byte(o.Y % 64),
	}
}
