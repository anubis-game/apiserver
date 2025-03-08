package vector

import (
	"math"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

func (v *Vector) shrink(hid byte) matrix.Coordinate {
	// Remember the current tail so we can use its value to shrink this Vector
	// below.

	cur := v.tai
	nxt := cur.nxt
	prv := cur.crd

	// Move the current tail node closer to the body instead of removing it
	// entirely, but only if the tail's neighbour node carries at least one hidden
	// node. We can calculate the last step of closing in on the body more
	// efficiently, because calculating the mid point of a line does not require
	// math.Sqrt().

	if nxt.hid == 1 {
		nxt.hid = 0
		cur.crd.X, cur.crd.Y = roundI(float64(cur.crd.X+nxt.crd.X)/2), roundI(float64(cur.crd.Y+nxt.crd.Y)/2)

		return prv
	} else if nxt.hid > 1 {
		nxt.hid -= int8(hid)
		cur.crd.X, cur.crd.Y = closer(cur.crd, nxt.crd)

		return prv
	}

	// The next item of the old tail becomes the new tail. Also reduce the
	// internal length counter.

	{
		v.tai = nxt
		v.tai.prv = nil
	}

	return prv
}

// closer brings "lef" closer to "rig" by returning the updated coordinates of
// "lef" so that it moves one standard distance towards "rig" on a 2 dimensional
// plane.
func closer(lef matrix.Coordinate, rig matrix.Coordinate) (int, int) {
	dsx, dsy := float64(rig.X-lef.X), float64(rig.Y-lef.Y)
	scl := nrm / math.Sqrt((dsx*dsx)+(dsy*dsy))
	return lef.X + roundI(dsx*scl), lef.Y + roundI(dsy*scl)
}
