package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Occupy returns those partition coordinates that are currently occupied by
// this Vector's body parts. Iterating over the list returned here and providing
// every item to Vector.Bounds() will return all segments and their respective
// partitions.
func (v *Vector) Occupy() []object.Object {
	return v.bpb
}
