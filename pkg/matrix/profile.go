package matrix

const (
	// SZ is the index for the size byte. This number describes the radius of a
	// player's head.
	SZ int = 0
	// TP is the index for the type byte.
	TP int = 1
)

// Profile contains the size and type bytes, describing the object's character
// and its evolution.
type Profile [2]byte
