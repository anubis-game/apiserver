package matrix

const (
	// QD is the index for the quadrant byte.
	QD int = 0
	// AG is the index for the angle byte.
	AG int = 1
)

// Space contains the quadrant and angle bytes, describing the direction in
// which a player is currently moving.
type Space [2]byte
