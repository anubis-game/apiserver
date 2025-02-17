package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/object"
)

// Expand moves the vector along the direction of the given target object and
// expands the underlying segments. After calling Expand, the underlying vector
// has 1 more object, which is the added target head.
func (v *Vector) Expand(hea object.Object) {
	lin := &Linker{val: hea}

	{
		v.hea.nxt = lin
		v.hea = lin
		v.len++
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

	// Always keep track of the amount of coordinates that occupy their respective
	// partitions.

	{
		v.xfr[prt.X]++
		v.yfr[prt.Y]++
	}

	if ind == 0 {
		// Initialize a new partition buffer with the Vector specific ID bytes, and
		// the given header bytes.

		{
			buf := make([]byte, 2+object.Len)

			copy(buf[:2], v.uid[:])
			copy(buf[2:], byt[:])

			v.buf[prt] = buf
		}

		// Only if the new header breaks into an unoccupied partition, only then do
		// we have to check in which direction we are overflowing. And then, update
		// our boundaries according to their direction of change.

		{
			v.occ.New = prt
		}

		if prt.Y > v.occ.Top {
			{
				v.occ.Top = prt.Y
			}

			for x := v.scr.Lef; x <= v.scr.Rig; x += matrix.Prt {
				v.scr.Prt = append(v.scr.Prt, object.Object{X: x, Y: v.scr.Top})
			}
		}
		if prt.X > v.occ.Rig {
			{
				v.occ.Rig = prt.X
			}

			for y := v.scr.Bot; y <= v.scr.Top; y += matrix.Prt {
				v.scr.Prt = append(v.scr.Prt, object.Object{X: v.scr.Rig, Y: y})
			}
		}
		if prt.Y < v.occ.Bot {
			{
				v.occ.Bot = prt.Y
			}

			for x := v.scr.Lef; x <= v.scr.Rig; x += matrix.Prt {
				v.scr.Prt = append(v.scr.Prt, object.Object{X: x, Y: v.scr.Bot})
			}
		}
		if prt.X < v.occ.Lef {
			{
				v.occ.Lef = prt.X
			}

			for y := v.scr.Bot; y <= v.scr.Top; y += matrix.Prt {
				v.scr.Prt = append(v.scr.Prt, object.Object{X: v.scr.Lef, Y: y})
			}
		}
	} else {
		// Extend the current buffer with the compressed 6 byte version of the given
		// coordinates. Using preallocated slices via copy safes about 5 ns/op
		// compared to using append. Note that using a new preallocated byte slice
		// fixes the memory leak incurred during Vector.shrink() where we merely
		// reslice the partition buffer.

		app := make([]byte, ind+object.Len)

		copy(app[:ind], buf)    // existing buffer goes first
		copy(app[ind:], byt[:]) // append the new header bytes

		v.buf[prt] = app

		// Reset the newly occupied partition every time the occupation is in fact
		// not new anymore.

		{
			v.occ.New = object.Object{}
		}
	}
}
