package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/object"
)

func (v *Vector) Shrink() object.Object {
	// Remember the current tail so we can use its value to shrink this Vector
	// below.

	tai := v.tai

	// The next item of the old tail becomes the new tail. Also reduce the
	// internal length counter.

	{
		v.tai = tai.nxt
		v.len--
	}

	// Shrink this Vector using its old tail value.

	{
		v.shrink(tai.val)
	}

	return tai.val
}

func (v *Vector) shrink(old object.Object) {
	prt := old.Prt()

	// Always keep track of the amount of coordinates that do not occupy their
	// respective partitions anymore.

	{
		v.xfr[prt.X]--
		v.yfr[prt.Y]--
	}

	// Reduce or delete the partition coordinates according to the current state.

	if len(v.occ.Prt[prt]) == 1 {
		// There is only one item left. That item is the object we are asked to
		// delete.

		{
			delete(v.occ.Prt, prt)
		}

		// Shrink the partition boundaries according to the direction of change as
		// specified by the old tail coordinates.

		tai := v.tai.val.Prt()

		if prt.Y == v.occ.Top && v.yfr[prt.Y] == 0 {
			{
				delete(v.yfr, prt.Y)
			}

			for x := v.scr.Lef; x <= v.scr.Rig; x += matrix.Prt {
				delete(v.scr.Prt, object.Object{X: x, Y: v.scr.Top})
			}

			{
				v.occ.Top = tai.Y
			}
		}

		if prt.X == v.occ.Rig && v.xfr[prt.X] == 0 {
			{
				delete(v.xfr, prt.X)
			}

			for y := v.scr.Bot; y <= v.scr.Top; y += matrix.Prt {
				delete(v.scr.Prt, object.Object{X: v.scr.Rig, Y: y})
			}

			{
				v.occ.Rig = tai.X
			}
		}

		if prt.Y == v.occ.Bot && v.yfr[prt.Y] == 0 {
			{
				delete(v.yfr, prt.Y)
			}

			for x := v.scr.Lef; x <= v.scr.Rig; x += matrix.Prt {
				delete(v.scr.Prt, object.Object{X: x, Y: v.scr.Bot})
			}

			{
				v.occ.Bot = tai.Y
			}
		}

		if prt.X == v.occ.Lef && v.xfr[prt.X] == 0 {
			{
				delete(v.xfr, prt.X)
			}

			for y := v.scr.Bot; y <= v.scr.Top; y += matrix.Prt {
				v.scr.Prt[object.Object{X: v.scr.Lef, Y: y}] = struct{}{}
			}

			{
				v.occ.Lef = tai.X
			}
		}
	} else {
		// The coordinates to remove are always represented by the very first item
		// in the coordinate partition.

		{
			v.occ.Prt[prt] = v.occ.Prt[prt][1:]
		}
	}
}
