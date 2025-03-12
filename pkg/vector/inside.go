package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// Inside returns the occupied coordinates of this Vector that the screen
// boundaries as returned by Vector.Screen() can see. If the given Vector nodes
// and screen boundaries do not overlap at all, then nil is returned. Nil is
// also returned if the given areas overlap, while no Vector coordinates are
// actually located in the overlapping area. The preliminary overlap is verified
// using an AABB check (Axis-Aligned Bounding Box).
func (v *Vector) Inside(stp int, srg int, sbt int, slf int) []matrix.Coordinate {
	// Check whether any overlap exists before attempting to collect any of the
	// overlapping partitions.

	if stp < v.obt || srg < v.olf || sbt > v.otp || slf > v.org {
		return nil
	}

	top := minInt(stp, v.otp)
	rig := minInt(srg, v.org)
	bot := maxInt(sbt, v.obt)
	lef := maxInt(slf, v.olf)

	var ins []matrix.Coordinate

	// Walk along all Vector nodes. We can afford the full loop here for several
	// reasons.
	//
	//     1. The number of partitions to iterate for any potential
	//        overlapping area is relatively small.
	//
	//     2. The number of Vector nodes within any given partition is
	//        relatively limited.
	//

	for n := v.hea; n != nil; n = n.prv {
		// If this node is not inside the overlapping area, then skip it.

		var p matrix.Partition
		{
			p = n.crd.Pt1()
		}

		if top < p.Y || rig < p.X || bot > p.Y || lef > p.X {
			continue
		}

		{
			ins = append(ins, n.crd)
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
