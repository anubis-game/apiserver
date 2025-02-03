package matrix

//     (123*4096)+(101*64)+117

const (
	// Max is the upper boundary of the game map in pixels, 64^3.
	Max = 262_144
	// Min is the upper boundary of the game map in pixels.
	Min = 0
	// Thr is the pixel threshold around the edges of the game map in which
	// players cannot be placed initially upon joining the game. The purpose of
	// this buffer region is to not put players too close to the edges of the
	// game, so that they can not run into the wall accidentally.
	Thr = 1_000
)
