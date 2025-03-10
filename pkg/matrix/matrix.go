package matrix

const (
	// Max is the upper boundary of the game map in pixels, 262,144.
	Max = 64 * 64 * 64
	// Min is the upper boundary of the game map in pixels.
	Min = 0
	// Prt is the scaling value allowing us to partition the X and Y axis into
	// logical buckets similar to quadtrees. 128*2048 = 64^3
	Prt = 128 // TODO:infra rename to Pt1 and add Pt8 with 1024 pixels for secondary partitions
	// Thr is the pixel threshold around the edges of the game map in which
	// players cannot be placed initially upon joining the game. The purpose of
	// this buffer region is to not put players too close to the edges of the
	// game, so that they can not run into the wall accidentally.
	Thr = 1_024
)
