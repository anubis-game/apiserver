package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

// Impact returns whether this Vector's head node collides with the provided
// impact node coordinate. The geometric check underneath is simply based on
// circle intersection.
func (v *Vector) Impact(inc matrix.Coordinate, inr byte) bool {
	// Normalize the two points to calculate the hypothenuse of a right triangle.

	var x float64
	var y float64
	{
		x = float64(v.hea.crd.X - inc.X)
		y = float64(v.hea.crd.Y - inc.Y)
	}

	// Sum the radii of both Vector nodes, to get the maximum allowed distance
	// between two players. What we effectively try to figure out here is whether
	// two circles overlap.

	var s float64
	{
		s = float64(v.crx.Rad) + float64(inr)
	}

	// Calculate the distance between two points using the normalized side lengths
	// of a right triangle. Only if the sum of the radii is larger than the
	// distance between the two nodes, only then do we have an intersection.

	return (s * s) >= ((x * x) + (y * y))
}
