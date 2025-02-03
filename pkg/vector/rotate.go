package vector

// Rotate moves the vector along the direction of the given target object
// without expanding the underlying amount of segments. After calling Rotate,
// the underlying vector has the same amount of objects as it had before.
func (v *Vector) Rotate(trg Object) {
	copy(v.obj[1:], v.obj[:v.len-1]) // shift without tail
	v.obj[0] = trg                   // target becomes head
}
