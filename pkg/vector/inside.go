package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// Inside returns the occupied coordinates of this Vector that the screen
// boundaries as returned by Vector.Screen() can see. If the given Vector nodes
// and screen boundaries do not overlap at all, then nil is returned. Nil is
// also returned if the given areas overlap, while no Vector coordinates are
// actually located in the overlapping area. The preliminary overlap is verified
// using an AABB check (Axis-Aligned Bounding Box). Vector.Inside() is only
// called once for new players joining the game, in order to get a full
// representation of a Vector's occupied coordinates. All other players will get
// reconciled based on the game map delta as provided by Vector.Change().
func (v *Vector) Inside(stp int, srg int, sbt int, slf int) map[matrix.Partition][]matrix.Coordinate {
	// Check whether any overlap exists before attempting to collect any of the
	// overlapping partitions.

	if srg < v.olf || stp < v.obt || slf > v.org || sbt > v.otp {
		return nil
	}

	top := minInt(stp, v.otp)
	rig := minInt(srg, v.org)
	bot := maxInt(sbt, v.obt)
	lef := maxInt(slf, v.olf)

	// TODO:test we need to verify that the correct coordinates can be found
	// within the correct partitions across various scenarious.

	var ins map[matrix.Partition][]matrix.Coordinate

	// Walk along all Vector nodes. We can afford the full loop here for several
	// reasons.
	//
	//     1. Vector.Inside() is called relatively rarely.
	//
	//     2. The number of partitions to iterate for any potential
	//        overlapping area is relatively small.
	//
	//     3. The number of Vector nodes within any given partition is
	//        relatively limited.
	//

	for n := v.hea; n != nil; n = n.prv {
		// If this node is not inside the overlapping area, then skip it.

		if rig < n.crd.X || top < n.crd.Y || lef > n.crd.X || bot > n.crd.Y {
			continue
		}

		// Only allocate a map if there are any Vector coordinates within the
		// verified overlap. It might be that rectancles overlap on one side,
		// while the Vector coordinates forming the rectangle are located on the
		// other side of the overlapping area. In such a case we iterate over all
		// overlapping partitions without finding any Vector coordinates. And so
		// we do not have to allocate any map before we really need it.

		if ins == nil {
			ins = map[matrix.Partition][]matrix.Coordinate{}
		}

		{
			ins[n.crd.Prt()] = append(ins[n.crd.Prt()], n.crd)
		}
	}

	return ins
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
