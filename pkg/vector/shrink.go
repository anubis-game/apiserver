package vector

import (
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
	buf := v.buf[prt]
	ind := len(buf)

	// Always keep track of the amount of coordinates that do not occupy their
	// respective partitions anymore.

	{
		v.xfr[prt.X]--
		v.yfr[prt.Y]--
	}

	// Reduce or delete the fanout buffer as described below.

	if ind == 3+object.Len {
		// There is only one item left. That item is the object we are asked to
		// delete.

		{
			delete(v.buf, prt)
		}

		// Shrink the partition boundaries according to the direction of change as
		// specified by the old tail coordinates.

		tai := v.tai.val.Prt()

		{
			v.occ.Old = prt
		}

		if prt.Y == v.occ.Top && v.yfr[prt.Y] == 0 {
			{
				v.occ.Top = tai.Y
			}

			{
				delete(v.yfr, prt.Y)
			}
		}
		if prt.X == v.occ.Rig && v.xfr[prt.X] == 0 {
			{
				v.occ.Rig = tai.X
			}

			{
				delete(v.xfr, prt.X)
			}
		}
		if prt.Y == v.occ.Bot && v.yfr[prt.Y] == 0 {
			v.occ.Bot = tai.Y

			{
				delete(v.yfr, prt.Y)
			}
		}
		if prt.X == v.occ.Lef && v.xfr[prt.X] == 0 {
			{
				v.occ.Lef = tai.X
			}

			{
				delete(v.xfr, prt.X)
			}
		}
	} else {
		// The item to remove is always represented by the very first 6 bytes of the
		// buffer after the 2 ID bytes. Note that we are only reslicing the existing
		// partition buffer, without deleting the remaining tail still allocated in
		// the underlying data array. This alone would usually imply a memory leak,
		// but we are fixing this memory leak in due time, once the partition buffer
		// gets deleted entirely as soon as the Vector moves out of it naturally
		// throughout the game.

		{
			buf[2]--                          // decrement coordinate amount
			copy(buf[3:], buf[3+object.Len:]) // remove the first coordinate
			v.buf[prt] = buf[:ind-object.Len] // cut the outdated buffer end
		}

		// Reset the recently occupied partition every time the occupation is in
		// fact not recent anymore.

		{
			v.occ.Old = object.Object{}
		}
	}
}
