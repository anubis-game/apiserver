package vector

// Inside returns whether this Vector resides inside the provided Screen. The
// check is done using a negative AABB check (Axis-Aligned Bounding Box).
func (v *Vector) Inside(scr *Screen) bool {
	return !(scr.Rig < v.occ.Lef || scr.Top < v.occ.Bot || scr.Lef > v.occ.Rig || scr.Bot > v.occ.Top)
}
