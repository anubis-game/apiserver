package vector

import "github.com/anubis-game/apiserver/pkg/object"

func (v *Vector) shrink(old object.Object) {
	prt := old.Prt()
	buf := v.buf[prt]
	siz := len(buf)

	// Reduce the fanout buffer given any of the situations described below.

	if siz == object.Len {

		// There is only one item left. That item is the object we are asked to
		// delete.

		{
			delete(v.buf, prt)
		}
	} else {

		// The item to remove is always the very first part of the buffer.

		{
			v.buf[prt] = buf[object.Len:]
		}
	}

	// Shrink the partition boundaries in case the removed coordinates are the
	// last on that edge.

	{
		v.xfr[prt.X]--
		v.yfr[prt.Y]--
	}

	if prt.Y == v.top && v.yfr[prt.Y] == 0 {
		v.top = v.tai.val.Prt().Y
	}
	if prt.X == v.rig && v.xfr[prt.X] == 0 {
		v.rig = v.tai.val.Prt().X
	}
	if prt.Y == v.bot && v.yfr[prt.Y] == 0 {
		v.bot = v.tai.val.Prt().Y
	}
	if prt.X == v.lef && v.xfr[prt.X] == 0 {
		v.lef = v.tai.val.Prt().X
	}

	if v.xfr[prt.X] == 0 {
		delete(v.xfr, prt.X)
	}
	if v.yfr[prt.Y] == 0 {
		delete(v.yfr, prt.Y)
	}
}
