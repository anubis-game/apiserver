package vector

import (
	"math"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func (v *Vector) shrink(hid byte) matrix.Coordinate {
	// Move the current tail closer to the body instead of removing it entirely,
	// but only if the tail's neighbour carry's at least one hidden segment.

	if v.tai.nxt.hid > 0 {
		old := v.tai.crd

		v.tai.nxt.hid -= int8(hid)
		v.tai.crd.X, v.tai.crd.Y = closer(v.tai.crd, v.tai.nxt.crd)

		return old
	}

	// Remember the current tail so we can use its value to shrink this Vector
	// below.

	old := v.tai

	// The next item of the old tail becomes the new tail. Also reduce the
	// internal length counter.

	{
		v.tai = old.nxt
		v.tai.prv = nil
	}

	return old.crd
}

// closer brings "lef" closer to "rig" by returning the updated coordinates of
// "lef" so that it moves one standard distance towards "rig" on a 2 dimensional
// plane.
func closer(lef matrix.Coordinate, rig matrix.Coordinate) (int, int) {
	dsx, dsy := float64(rig.X-lef.X), float64(rig.Y-lef.Y)
	scl := nrm / math.Sqrt((dsx*dsx)+(dsy*dsy))
	return lef.X + roundI(dsx*scl), lef.Y + roundI(dsy*scl)
}
