package vector

// hidden returns the amount of hidden segments that we create by modifying the
// head and tail coordinates instead of replacing them on a per frame basis.
func (v *Vector) hidden() int {
	var hid int

	cur := v.hea
	for cur != nil {
		hid += int(cur.hid)
		cur = cur.prv
	}

	return hid
}
