package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Expand moves the vector along the direction of the given target object and
// expands the underlying segments. After calling Expand, the underlying vector
// has 1 more object, which is the added target head.
func (v *Vector) Expand(hea object.Object) {
	lin := &Linker{val: hea}

	{
		v.hea.nxt = lin
		v.hea = lin
		v.siz++
	}

	{
		v.expand(hea)
	}
}

func (v *Vector) expand(hea object.Object) {
	prt := hea.Prt()
	buf := v.buf[prt]
	byt := hea.Byt()
	ind := len(buf)

	// Extend the current buffer with the compressed 6 byte version of the given
	// coordinates. Using preallocated slices via copy gives us 5 ns/op as
	// compared to using append. Note that using a new preallocated byte slice
	// fixes the memory leak incurred during Vector.shrink() where we merely
	// reslice the partition buffer.

	if ind == 0 {
		v.buf[prt] = byt[:]
	} else {
		ext := make([]byte, ind+object.Len)

		copy(ext[:ind], buf)
		copy(ext[ind:], byt[:])

		v.buf[prt] = ext
	}

	// In case any of the given partition overflows our currently known
	// boundaries, update those boundaries according to their direction of change.

	{
		v.xfr[prt.X]++
		v.yfr[prt.Y]++
	}

	if prt.Y > v.top {
		v.top = prt.Y
	}
	if prt.Y < v.bot {
		v.bot = prt.Y
	}
	if prt.X > v.rig {
		v.rig = prt.X
	}
	if prt.X < v.lef {
		v.lef = prt.X
	}
}
