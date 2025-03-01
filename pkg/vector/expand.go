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

	// Always keep track of the amount of coordinates that occupy their respective
	// partitions.

	{
		v.xfr[prt.X]++
		v.yfr[prt.Y]++
	}

	if len(v.occ.Prt[prt]) == 0 {
		// Initialize a new partition with the Vector head.

		{
			v.occ.Prt[prt] = []object.Object{hea}
		}

		// Only if the new header breaks into an unoccupied partition, only then do
		// we have to check in which direction we are overflowing. And then, update
		// our boundaries according to their direction of change.

		if prt.Y > v.occ.Top {
			{
				v.occ.Top = prt.Y
			}

			for x := v.scr.Lef; x <= v.scr.Rig; x += matrix.Prt {
				v.scr.Prt[object.Object{X: x, Y: v.scr.Top}] = struct{}{}
			}
		}

		if prt.X > v.occ.Rig {
			{
				v.occ.Rig = prt.X
			}

			for y := v.scr.Bot; y <= v.scr.Top; y += matrix.Prt {
				v.scr.Prt[object.Object{X: v.scr.Rig, Y: y}] = struct{}{}
			}
		}

		if prt.Y < v.occ.Bot {
			{
				v.occ.Bot = prt.Y
			}

			for x := v.scr.Lef; x <= v.scr.Rig; x += matrix.Prt {
				v.scr.Prt[object.Object{X: x, Y: v.scr.Bot}] = struct{}{}
			}
		}

		if prt.X < v.occ.Lef {
			{
				v.occ.Lef = prt.X
			}

			for y := v.scr.Bot; y <= v.scr.Top; y += matrix.Prt {
				v.scr.Prt[object.Object{X: v.scr.Lef, Y: y}] = struct{}{}
			}
		}
	} else {
		// Extend the existing partition with the new Vector head.

		{
			v.occ.Prt[prt] = append(v.occ.Prt[prt], hea)
		}
	}
}
