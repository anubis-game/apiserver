package matrix

const (
	// Max is the upper byte boundary of any bucket and pixel. Buckets and pixels
	// above this level overflow.
	Max int = 163
	// Min is the lower byte boundary of any bucket and pixel. Buckets and pixels
	// below this level underflow.
	Min int = 100
	// Siz is the bucket size of the layered coordinate system. This is the total
	// amount of outer buckets within the entire coordinate system. This is also
	// the number of inner buckets within any given outer bucket. And this is the
	// quadratic length in pixels per inner bucket. Player positions may overflow
	// and underflow into other buckets, if players move beyond the boundaries of
	// their current position.
	Siz int = Max - Min + 1
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

// decByt decrements the bytes of any given axis of the layered coordinate
// system and manages eventual underflows accordingly.
func decByt(out byte, inn byte) (byte, byte) {
	if inn == byte(Min) {
		return out - 1, byte(Max)
	}

	return out, inn - 1
}

// incByt increments the bytes of any given axis of the layered coordinate
// system and manages eventual overflows accordingly.
func incByt(out byte, inn byte) (byte, byte) {
	if inn == byte(Max) {
		return out + 1, byte(Min)
	}

	return out, inn + 1
}
