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

		// Initialize a new partition buffer with the given header bytes.

		{
			v.buf[prt] = byt[:]
		}

		// Only if the new header breaks into an unoccupied partition, only then do
		// we have to check in which direction we are overflowing. And then, update
		// our boundaries according to their direction of change.

		{
			v.vpb = nil
		}

		if prt.Y > v.btp {
			{
				v.btp = prt.Y
			}

			for i := v.vlf; i <= v.vrg; i += matrix.Prt {
				v.vpb = append(v.vpb, object.Object{X: i, Y: v.vtp})
			}
		}
		if prt.X > v.brg {
			{
				v.brg = prt.X
			}

			for i := v.vbt; i <= v.vtp; i += matrix.Prt {
				v.vpb = append(v.vpb, object.Object{X: v.vrg, Y: i})
			}
		}
		if prt.Y < v.bbt {
			{
				v.bbt = prt.Y
			}

			for i := v.vlf; i <= v.vrg; i += matrix.Prt {
				v.vpb = append(v.vpb, object.Object{X: i, Y: v.vbt})
			}
		}
		if prt.X < v.blf {
			{
				v.blf = prt.X
			}

			for i := v.vbt; i <= v.vtp; i += matrix.Prt {
				v.vpb = append(v.vpb, object.Object{X: v.vlf, Y: i})
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
	}
}
