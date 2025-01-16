package coordinate

import (
	"math"
)

const (
	// bmx is the maximum bucket index of the layered coordinate system.
	bmx float64 = 31
	// bsz is the bucket size of the layered coordinate system. This is the total
	// amount of outer buckets within the entire coordinate system. This is also
	// the number of inner buckets within any given outer bucket. And this is the
	// quadratic length in pixels per inner bucket. Player positions may overflow
	// and underflow into other buckets, if players move beyond the boundaries of
	// their current position.
	bsz float64 = bmx + 1
	// dis is the standard distance travelled in pixels per millisecond. For
	// instance, 0.192 px/ms implies 192 px/s, which equates to 6 buckets at 32
	// pixels per bucket.
	dis float64 = 0.192
	// qrd is the quadrant specific radian factor for half Pi. This is the atomic
	// amount of radians applied to a single byte of the quadrant specific angle
	// spc[1]. Multiplying qrd by the angle byte spc[1] provides the radians to
	// calculate a player's coordinate displacement most efficiently.
	//
	//     (spc[1] / 255) * (Pi / 2)
	//     (spc[1] / 255) * 1.570796
	//     spc[1] * (1 / 255 * 1.570796)
	//     spc[1] * 0.006159984314
	//
	qrd float64 = 0.006159984314
)

var (
	// sin is the sine lookup table to cache all possible sine values based on any
	// given angle byte.
	sin [256]float64
	// cos is the cosine lookup table to cache all possible cosine values based on
	// any given angle byte.
	cos [256]float64
)

func init() {
	for i := 0; i < 256; i++ {
		var rad float64
		{
			rad = float64(i) * qrd
		}

		{
			sin[i] = math.Sin(rad)
			cos[i] = math.Cos(rad)
		}
	}
}

func Next(cur [6]byte, spc [2]byte, tim [2]byte) ([6]byte, byte) {
	// cur is the current possition of a player expressed in a layered coordinate
	// system. The first byte pair x0 and y0 defines the outer buckets that the
	// entire game map consists of. The second byte pair x1 and y1 defines the
	// inner bucket within the associated outer bucket. The third byte pair x2 and
	// y2 defines the position pixels within the referenced inner bucket.
	//
	//     [
	//       x0, y0,
	//       x1, y1,
	//       x2, y2,
	//     ]
	//

	var x0, y0 float64
	var x1, y1 float64
	var x2, y2 float64
	{
		x0, y0 = float64(cur[0]), float64(cur[1])
		x1, y1 = float64(cur[2]), float64(cur[3])
		x2, y2 = float64(cur[4]), float64(cur[5])
	}

	// tim contains the time bytes including a millisecond duration and a velocity
	// factor. The elapsed duration tim[0] contains the byte encoded milliseconds
	// that passed between the previous and the current update cycle of movement.
	// This delta can be imagined as the elapsed time between the previous
	// timestamp A and the current timestamp B, during which a player was moving
	// through the game. The velocity factor tim[1] describes at which speed a
	// player is moving across the field. The standard velocity is 0x01, or 100%.
	// E.g. an accelerated velocity of 400% would be encoded as 0x04.
	//
	//     time under velocity
	//
	//     ----A--------B---->
	//

	var tot float64
	{
		tot = dis * float64(tim[0]) * float64(tim[1])
	}

	// spc contains the space bytes including a quadrant indicator and the angle
	// alpha. In a coordinate system of 4 quadrants, spc[0] is one of [0x01, 0x02,
	// 0x03, 0x04], indicating one of the quadrants towards which a player is
	// moving right now. Given any quadrant, alpha is the clock wise angle encoded
	// as a single byte in the range of [0, 255], dividing 90 degrees of any
	// quadrant into 256 possible angles. The measurement of alpha starts at 0°
	// for quadrant 1, 90° for quadrant 2, 180° for quadrant 3, and 270° for
	// quadrant 4.
	//
	//                       0°
	//
	//             +---------+---------+
	//             |         |         |
	//             |    4    |    1    |
	//             |         |         |
	//     270°    +---------+---------+    90°
	//             |         |         |
	//             |    3    |    2    |
	//             |         |         |
	//             +---------+---------+
	//
	//                      180°
	//

	switch spc[0] {
	case 0x01:
		x2 += math.Round(tot * sin[spc[1]])
		y2 += math.Round(tot * cos[spc[1]])
	case 0x02:
		x2 += math.Round(tot * cos[spc[1]])
		y2 -= math.Round(tot * sin[spc[1]])
	case 0x03:
		x2 -= math.Round(tot * sin[spc[1]])
		y2 -= math.Round(tot * cos[spc[1]])
	case 0x04:
		x2 -= math.Round(tot * cos[spc[1]])
		y2 += math.Round(tot * sin[spc[1]])
	}

	// The calculated pixel movement may result in valid or invalid underflows and
	// overflows. The valid version of those boundary jumps implies to move
	// forward to the following inner and outer coordinate buckets. The invalid
	// version of the respective underflow and overflow violations appear on the
	// very edges of our coordinate system. Once we detect such a violation we
	// return the overflow bytes for the lower case letters 't', 'r', 'b' and 'l'.
	// Those boundary violation bytes represent the underflows and overflows
	// towards the top, right, bottom and left respectively.

	for x2 >= bsz {
		{
			x2 -= bsz
			x1++
		}

		if x1 >= bsz {
			x1 -= bsz
			x0++
		}

		if x0 >= bmx && x1 >= bmx {
			// Overflow to the right, beyond the allowed positive x-axis boundary.
			return [6]byte{}, byte('r')
		}
	}

	for x2 < 0 {
		if x0 <= 0 && x1 <= 0 {
			// Underflow to the left, beyond the allowed negative x-axis boundary.
			return [6]byte{}, byte('l')
		}

		{
			x2 += bsz
			x1--
		}

		if x1 < 0 {
			x1 += bsz
			x0--
		}
	}

	for y2 >= bsz {
		{
			y2 -= bsz
			y1++
		}

		if y1 >= bsz {
			y1 -= bsz
			y0++
		}

		if y0 >= bmx && y1 >= bmx {
			// Overflow to the top, beyond the allowed positive y-axis boundary.
			return [6]byte{}, byte('t')
		}
	}

	for y2 < 0 {
		if y0 <= 0 && y1 <= 0 {
			// Overflow to the bottom, beyond the allowed negative y-axis boundary.
			return [6]byte{}, byte('b')
		}

		{
			y2 += bsz
			y1--
		}

		if y1 < 0 {
			y1 += bsz
			y0--
		}
	}

	// We return the updated position, given the current position, the direction
	// of movement, and the velocity at which a player moves during a standard
	// frame duration. The returned position bytes rounded to floating point
	// precision, which means that we are only returning the nearest full pixel
	// changes of movement with mirror consistency. Code executing at this point
	// does not represent a boundary violation of our coordinate system. Therefore
	// we return the empty byte as second argument.

	return [6]byte{byte(x0), byte(y0), byte(x1), byte(y1), byte(x2), byte(y2)}, 0x00
}
