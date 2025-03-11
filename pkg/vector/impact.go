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

	// Sum the player sizes to get the maximum allowed distance between two
	// players. The size of each player is the radius byte around each point of
	// each player. So what we effectively try to figure out here is whether two
	// circles overlap.

	var s float64
	{
		s = float64(v.Charax().Rad) + float64(inr)
	}

	// Calculate the distance between two points using the normalized side lengths
	// of a right triangle. Only of the sum of the radiuses is larger than the
	// distance between the two points, only then do we have an intersection. The
	// optimization we are doing here is to prevent the more expensive downward
	// computation of square root, and replace it with the comparitively more
	// efficient equivalent upward computation, by squaring the sum of both
	// radiuses.

	return (s * s) >= ((x * x) + (y * y))
}
