package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// Inside executes the given callback for every occupied coordinate of this
// Vector that the screen boundaries as returned by Vector.Screen() can see. If
// the given Vector nodes and screen boundaries do not overlap at all, then fnc
// is not executed at all. The same is true if the given areas overlap, while no
// Vector coordinates are actually located in the overlapping area. The
// preliminary overlap is verified using an AABB check (Axis-Aligned Bounding
// Box). Inside stops processing if fnc returns false.
func (v *Vector) Inside(stp int, srg int, sbt int, slf int, fnc func(matrix.Coordinate) bool) {
	// Check whether any overlap exists before attempting to collect any of the
	// overlapping partitions.

	if stp < v.obt || srg < v.olf || sbt > v.otp || slf > v.org {
		return
	}

	// Define the boundaries of the overlapping area in terms of partition
	// coordinates.

	top := minInt(stp, v.otp)
	rig := minInt(srg, v.org)
	bot := maxInt(sbt, v.obt)
	lef := maxInt(slf, v.olf)

	// Walk along all Vector nodes. We can afford the full loop here for several
	// reasons.
	//
	//     1. The number of partitions to iterate for any potential
	//        overlapping area is relatively small.
	//
	//     2. The number of nodes within any given Vector is relatively limited.
	//
	//     3. We do not allocate by calling Vector.Inside().
	//

	for n := v.hea; n != nil; n = n.prv {
		var p matrix.Partition
		{
			p = n.crd.Pt1()
		}

		// Skip this node if it is not inside the overlapping partition.

		if top < p.Y || rig < p.X || bot > p.Y || lef > p.X {
			continue
		}

		if !fnc(n.crd) {
			break
		}
	}
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
