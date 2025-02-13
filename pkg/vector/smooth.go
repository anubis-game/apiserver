package vector

import (
	"math"

	"github.com/anubis-game/apiserver/pkg/object"
)

const (
	// Sf is the smoothing factor for reducing and shifting the altitude of an
	// apex point by shrinking its connecting sides within a triangle.
	Sf float64 = 0.2
)

// Smooth constricts this Vector in O(N-2), where N is the number of segments
// within the underlying linked list. Constriction works by shortening the sides
// that connect 3 points, causing the middle point to be adjusted inwards.
func (v *Vector) Smooth() {
	// Define the first set of points that we start out with. We begin at the
	// tail, because the linked list does only provide the next segment upwards.

	lef := v.tai
	mid := lef.nxt
	rig := mid.nxt

	// In order to prevent an additional "if" condition within the loop below, we
	// call smooth() the first time before the loop starts.

	prx, pry := smooth(lef.val, mid.val, rig.val)

	// Shift our working set of points one more time. We want the smoothing to
	// only reflect the current state of this Vector. So we remember the
	// coordinate changes that we compute as we iterate, and only apply them to
	// the Vector segments once the computed changes cannot affect other
	// computations done within this update cycle. This means we have to shift
	// twice before we can apply coordinate changes to the segment that then is
	// not part of the set of points anymore which we use to compute coordinate
	// updates.

	lef = lef.nxt
	mid = mid.nxt
	rig = rig.nxt

	// Iterate until we find the head element, which has no next item linked
	// anymore. Meaning, once we get to the head, "rig" becomes nil.

	for rig != nil {
		// Simply compute the updated coordinates for the mid point.

		crx, cry := smooth(lef.val, mid.val, rig.val)

		// Remember this mid point's updated coordinates, but only after we applied
		// the previously remembered changes to the out of range segment.

		lef.val.X, lef.val.Y = prx, pry
		prx, pry = crx, cry

		// Simply shift our set of points along the linked list.

		lef = lef.nxt
		mid = mid.nxt
		rig = rig.nxt
	}

	// Ensure to apply the last update that could not be finished within the loop
	// above. The segment getting its coordinates updated here is the item
	// pointing to the head of the linked list.

	{
		lef.val.X, lef.val.Y = prx, pry
	}
}

// smooth is used to lower the apex point "mid" within a triangle, connected to
// its neighbours "lef" and "rig". The smoothing happening here does not create
// rounder curves, but instead simply constricts a vector when applied
// iteratively to all points within said vector.
func smooth(lef object.Object, mid object.Object, rig object.Object) (int, int) {
	// We simply scale the distance between the given points using the constant
	// factor Sf.

	smx := math.Round((Sf * float64(lef.X-mid.X)) + (Sf * float64(rig.X-mid.X)))
	smy := math.Round((Sf * float64(lef.Y-mid.Y)) + (Sf * float64(rig.Y-mid.Y)))

	return mid.X + int(smx), mid.Y + int(smy)
}
