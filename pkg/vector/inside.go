package vector

// Inside returns whether this Vector resides inside the outer bounds of
// somebody's screen given a successful AABB check (Axis-Aligned Bounding Box).
// The order of boundaries provided must be top, right, bottom, left as returned
// by Vector.Screen().
func (v *Vector) Inside(top int, rig int, bot int, lef int) bool {
	return !(rig < v.blf || top < v.bbt || lef > v.brg || bot > v.btp)
}
