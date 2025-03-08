package vector

// Screen returns the outer partition coordinates of this Vector's current view.
// The returned boundaries describe top, right, bottom and left, in this order.
func (v *Vector) Screen() (int, int, int, int) {
	return v.stp, v.srg, v.sbt, v.slf
}
