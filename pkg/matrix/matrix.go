package matrix

const (
	// Max is the upper boundary of the game map in pixels, 64^3.
	Max = 262_144
	// Min is the upper boundary of the game map in pixels.
	Min = 0
	// Prt is the scaling value allowing us to partition the X and Y axis into
	// logical buckets similar to quadtrees. 512^2 = 64^3
	Prt = 512
	// Thr is the pixel threshold around the edges of the game map in which
	// players cannot be placed initially upon joining the game. The purpose of
	// this buffer region is to not put players too close to the edges of the
	// game, so that they can not run into the wall accidentally.
	Thr = 1_024
)
