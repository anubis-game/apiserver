package matrix

func Intersect(obc Bucket, opx Pixel, osz byte, ebc Bucket, epx Pixel, esz byte) bool {
	// Normalize the origin and enemy buckets and pixels to calculate the
	// hypothenuse of a right triangle.
	//
	//     for y, check which point is greater, y - 100 = top distance
	//     for x, check which point is greater, x - 100 = right distance
	//     for y, the other point is smaller,   163 - y = bottom distance
	//     for x, the other point is smaller,   163 - x = left distance
	//
	//     top + bottom = a
	//     left + right = b
	//

	var a float64
	var b float64
	{
		a = sidLen(obc, opx, ebc, epx, X0, X1, X2)
		b = sidLen(obc, opx, ebc, epx, Y0, Y1, Y2)
	}

	// Sum the player sizes to get the maximum allowed distance between two
	// players. The size of each player is the radius byte around each point of
	// each player. So what we effectively try to figure out here is whether two
	// circles overlap.

	var s float64
	{
		s = float64(osz + esz)
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

func sidLen(obc Bucket, opx Pixel, ebc Bucket, epx Pixel, c0 int, c1 int, c2 int) float64 {
	var oc0 byte
	var oc1 byte
	var oc2 byte
	var ec0 byte
	var ec1 byte
	var ec2 byte
	{
		oc0 = obc[c0]
		oc1 = obc[c1]
		oc2 = opx[c2]
		ec0 = ebc[c0]
		ec1 = ebc[c1]
		ec2 = epx[c2]
	}

	if oc0 == ec0 {
		if oc1 == ec1 {
			//
			//     outer and inner bucket equal
			//

			if oc2 > ec2 {
				return float64(oc2 - ec2)
			} else {
				return float64(ec2 - oc2)
			}
		} else {
			//
			//     outer bucket equal
			//

			{
				if oc1 > ec1 {
					//
					//     opx above
					//
					//     epx below
					//
					return float64((oc2 - byte(Min)) + (byte(Max) - ec2))
				} else {
					//
					//     epx above
					//
					//     opx below
					//
					return float64((ec2 - byte(Min)) + (byte(Max) - oc2))
				}
			}
		}
	}

	//
	//     outer bucket above
	//

	if oc0 > ec0 {
		if oc1 > ec1 {
			//
			//     opx above
			//
			//     epx below
			//
			return float64((oc2 - byte(Min)) + (byte(Max) - ec2))
		} else {
			//
			//     epx above
			//
			//     opx below
			//
			return float64((ec2 - byte(Min)) + (byte(Max) - oc2))
		}
	}

	//
	//     outer bucket below
	//

	{
		if oc1 < ec1 {
			//
			//     opx above
			//
			//     epx below
			//
			return float64((oc2 - byte(Min)) + (byte(Max) - ec2))
		} else {
			//
			//     epx above
			//
			//     opx below
			//
			return float64((ec2 - byte(Min)) + (byte(Max) - oc2))
		}
	}
}
