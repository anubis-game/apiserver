package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

const (
	// Sf is the smoothing factor for reducing and shifting the altitude of an
	// apex point by shrinking its connecting sides within a triangle. E.g. a
	// value of 0.5 implies an apex reduction of 50%.
	Sf float64 = 0.5
)

// Smooth constricts this Vector in O(N-2), where N is the number of segments
// within the underlying linked list. Constriction works by shortening the sides
// that connect 3 points, causing the middle point to be adjusted inwards.
func (v *Vector) Smooth() {
	// Define the first set of points that we start out with. We begin at the
	// head, because we want to pull the body parts towards teh direction of
	// travel. This mimics real physics more accurately than traversing backwards
	// from tail to head.

	rig := v.hea
	mid := rig.prv
	lef := mid.prv

	// Iterate until we find the tail segment, which has no previous segment
	// linked anymore. So once we get to the tail, then "lef" becomes nil.

	for lef != nil {
		// Simply compute the updated coordinates for the mid point, but only if there
		// is enough space between the given nodes.

		// TODO:infra smoothing reduces the distance between the nodes. Every node
		// may account for hidden nodes, which represents the distance to the
		// previous neigbour node.
		//
		//     1. We have to decrement the hidden count of every node that moves so
		//        close to its previous neighbour that the full length of a normal
		//        distance has vanished. E.g. node B and node A have been 4 normal
		//        distances apart originally, and due to continous smoothing have
		//        moved closer to one another so that the effective distance between
		//        node B and node A reduced to the length of 3 normal distances. In
		//        such a case we have to decrement the hidden count of node A.
		//
		//     2. We have to enforce a minimum distance between neigbouring nodes.
		//        This minimum is already enforced naturally via Vector expansion
		//        and shrinking, but smoothing reduces the distance between nodes
		//        iteratively. The distance between nodes should then also reflect
		//        the player's size, so that smaller players maintain nodes closer
		//        to one another than larger players.
		//

		if mid.hid > 1 && rig.hid > 1 {
			mid.crd.X, mid.crd.Y = smooth(lef.crd, mid.crd, rig.crd)
		}

		// Simply shift our set of points along the linked list towards the tail.

		rig = rig.prv
		mid = mid.prv
		lef = lef.prv
	}
}

// smooth is used to lower the apex point "mid" within a triangle, connected to
// its neighbours "lef" and "rig". The smoothing happening here does not create
// rounder curves, but instead simply constricts a vector when applied
// iteratively to all points within said vector. In our implementation we simply
// scale the distance between the given points using the constant factor Sf.
// Note that we round via integer truncation.
func smooth(lef matrix.Coordinate, mid matrix.Coordinate, rig matrix.Coordinate) (int, int) {
	return mid.X + roundI(((float64(lef.X+rig.X)/2)-float64(mid.X))*Sf), mid.Y + roundI(((float64(lef.Y+rig.Y)/2)-float64(mid.Y))*Sf)
}
