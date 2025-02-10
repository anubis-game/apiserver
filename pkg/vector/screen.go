package vector

// Screen returns the outer boundaries of this Vector's view, expressed in
// partition coordinates. The order of boundaries returned is top, right,
// bottom, left.
func (v *Vector) Screen() (int, int, int, int) {
	return v.vtp, v.vrg, v.vbt, v.vlf
}
