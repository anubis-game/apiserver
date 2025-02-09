package vector

// Bounds returns this Vector's outer boundaries, expressed in partition
// coordinates.
func (v *Vector) Bounds() (int, int, int, int) {
	return v.top, v.rig, v.bot, v.lef
}
