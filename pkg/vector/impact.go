package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

func Impact(oxy matrix.Coordinate, osz byte, txy matrix.Coordinate, tsz byte) bool {
	// Normalize the two points to calculate the hypothenuse of a right triangle.

	var a float64
	var b float64
	{
		a = absInt(oxy.X - txy.X)
		b = absInt(oxy.Y - txy.Y)
	}

	// Sum the player sizes to get the maximum allowed distance between two
	// players. The size of each player is the radius byte around each point of
	// each player. So what we effectively try to figure out here is whether two
	// circles overlap.

	var s float64
	{
		s = float64(osz) + float64(tsz)
	}

	// Calculate the distance between two points using the normalized side lengths
	// of a right triangle. Only of the sum of the radiuses is larger than the
	// distance between the two points, only then do we have an intersection. The
	// optimization we are doing here is to prevent the more expensive downward
	// computation of square root, and replace it with the comparitively more
	// efficient equivalent upward computation, by squaring the sum of both
	// radiuses.

	return (s * s) > ((a * a) + (b * b))
}

func absInt(i int) float64 {
	if i >= 0 {
		return float64(i)
	}

	return -float64(i)
}
