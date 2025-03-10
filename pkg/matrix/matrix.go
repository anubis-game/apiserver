package matrix

const (
	// Max is the upper boundary of the game map in pixels, 262,144.
	Max = 64 * 64 * 64
	// Min is the upper boundary of the game map in pixels.
	Min = 0
	// Pt1 is the side length of a small partition in pixels, dividing the game
	// map into smaller logical buckets. There are 2048*2048 small partitions
	// based on Max.
	Pt1 = 128
	// Pt8 is the side length of a large partition in pixels, dividing the game
	// map into larger logical buckets. There are 256*256 large partitions based
	// on Max.
	Pt8 = 128 * 8
)
