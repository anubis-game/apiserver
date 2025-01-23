package matrix

const (
	// Max is the upper byte boundary of any bucket and pixel. Buckets and pixels
	// above this level overflow.
	Max byte = 163
	// Min is the lower byte boundary of any bucket and pixel. Buckets and pixels
	// below this level underflow.
	Min byte = 100
	// Siz is the bucket size of the layered coordinate system. This is the total
	// amount of outer buckets within the entire coordinate system. This is also
	// the number of inner buckets within any given outer bucket. And this is the
	// quadratic length in pixels per inner bucket. Player positions may overflow
	// and underflow into other buckets, if players move beyond the boundaries of
	// their current position.
	Siz byte = Max - Min + 1
	// Dia is the truncated diagonal pixel length of every inner bucket, minus 1.
	// This is the maximum amount of movement allowed to happen within a single
	// frame. Any player movement beyond this pixel length is invalid, causing an
	// illegal overflow.
	//
	//     ( Siz * math.Sqrt2 ) - 1
	//
	Dia float64 = 88
)

type Bucket [4]byte

func (b Bucket) Scale(siz byte) Bucket {
	del := Max - siz

	sx0, sy0, sx1, sy1 := b[X0], b[Y0], b[X1], b[Y1]

	if del < sx1 {
		sx0, sx1 = sx0+1, Min+(sx1-del)-1
	} else {
		sx1 += siz
	}

	if del < sy1 {
		sy0, sy1 = sy0+1, Min+(sy1-del)-1
	} else {
		sy1 += siz
	}

	return Bucket{sx0, sy0, sx1, sy1}
}
