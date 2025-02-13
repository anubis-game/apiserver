package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Bounds returns the most recently discovered partition coordinates. For a new
// Vector those bounds represent the Vector's entire range of sight. For Vectors
// being expanded or rotated those bounds represent the slice of partition
// coordinates on the screen that have just been revealed by movement towards
// any given direction.
func (v *Vector) Bounds() []object.Object {
	return v.vpb
}
