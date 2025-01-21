package matrix

import (
	"math"
)

const (
	// dms is the standard distance travelled in pixels per millisecond at a speed
	// of 100%. For instance, 0.192 px/ms implies 192 px/s, which equates to 3
	// buckets at 64 pixels per bucket.
	dms float64 = 0.192
	// qrf is the quadrant specific radian factor for half Pi. This is the atomic
	// amount of radians applied to a single byte of the quadrant specific angle
	// spc[1]. Multiplying qrf by the angle byte spc[1] provides the radians to
	// calculate a player's coordinate displacement most efficiently.
	//
	//     (spc[1] / 255) * (Pi / 2)
	//     (spc[1] / 255) * 1.570796
	//     spc[1] * (1 / 255 * 1.570796)
	//     spc[1] * 0.006159984314
	//
	qrf float64 = 0.006159984314
)

var (
	// cos is the cosine lookup table to cache all possible cosine values based on
	// any given angle byte.
	//
	//     cos[spc[1]]
	//
	cos [256]float64
	// sin is the sine lookup table to cache all possible sine values based on any
	// given angle byte.
	//
	//     sin[spc[1]]
	//
	sin [256]float64
	// dis is the distance lookup table to cache all possible time values based on
	// any given time byte.
	//
	//     dis[tim[0]][tim[1]]
	//
	dis [256][256]float64
)

func init() {
	// spc[0] and tim[0]
	for i := 0; i < 256; i++ {
		{
			cos[i] = math.Cos(float64(i) * qrf)
			sin[i] = math.Sin(float64(i) * qrf)
		}

		// tim[1]
		for j := 0; j < 256; j++ {
			dis[i][j] = dms * float64(i) * float64(j)
		}
	}
}

// Target uses the given origin to calculate the next point in a two dimensional
// coordinate system. Target effectively defines a line between the provided
// origin and the computed target. Both origin and target are represented here
// as Bucket key and Pixel location. Target does not support player movements of
// multiple buckets and may return a non zero overflow byte as third return
// value.
func Target(obc Bucket, opx Pixel, spc [2]byte, tim [2]byte) (Bucket, Pixel, byte) {
	// obc is the origin bucket and opx is the origin pixel. The origin describes
	// the current possition of a player within a layered coordinate system. The
	// first byte pair x0 and y0 refers to the outer buckets that the entire game
	// map consists of. The second byte pair x1 and y1 refers to the inner bucket
	// within the associated outer bucket. The third byte pair x2 and y2 refers to
	// the pixel location within the referenced inner bucket.
	//
	//     [
	//       x0, y0,    outer bucket
	//       x1, y1,    inner bucket
	//     ]
	//
	//     [
	//       x2, y2,    pixel location
	//     ]
	//

	var tx0 byte
	var ty0 byte
	var tx1 byte
	var ty1 byte
	var tx2 byte
	var ty2 byte
	{
		tx0 = obc[X0]
		ty0 = obc[Y0]
		tx1 = obc[X1]
		ty1 = obc[Y1]
		tx2 = opx[X2]
		ty2 = opx[Y2]
	}

	// tim contains the time bytes including a millisecond duration and a velocity
	// factor. The elapsed duration tim[0] contains the byte encoded milliseconds
	// that passed between the previous and the current update cycle of movement.
	// This delta can be imagined as the elapsed time between the previous
	// timestamp A and the current timestamp B, during which a player was moving
	// through the game. The velocity factor tim[1] describes at which speed a
	// player is moving across the field. The standard velocity is 0x01, or 100%.
	// E.g. an accelerated velocity of 400% would be encoded as 0x04. Velocity
	// factors beyond a certain threshold may cause player movements across
	// multiple buckets. So if the total distance travelled ends up being higher
	// than the maximum allowed travel distance, then we return a generic overflow
	// byte.
	//
	//     time under velocity
	//
	//     ----A--------B---->
	//

	var tot float64
	{
		tot = dis[tim[0]][tim[1]]
	}

	if tot > Dia {
		return Bucket{}, Pixel{}, byte('o')
	}

	// spc contains the space bytes including a quadrant indicator and the angle
	// alpha. In a coordinate system of 4 quadrants, spc[0] is one of [0x01, 0x02,
	// 0x03, 0x04], indicating one of the quadrants towards which a player is
	// moving right now. Given any quadrant, alpha is the clock wise angle encoded
	// as a single byte in the range of [0, 255], dividing 90 degrees of any
	// quadrant into 256 possible angles. The measurement of alpha starts at 0°
	// for quadrant 1, 90° for quadrant 2, 180° for quadrant 3, and 270° for
	// quadrant 4. The distance travelled from one point to another is given as
	// absolute uint8, calculated precisely as float64, and then rounded
	// efficiently via integer truncation by adding 0.5 to the computed delta. All
	// we have to do now in order to get to the next point is to add or remove the
	// integer distance to and from the x2 and y2 coordinates.
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

	var dc byte
	var ds byte
	{
		dc = byte(tot*cos[spc[1]] + 0.5)
		ds = byte(tot*sin[spc[1]] + 0.5)
	}

	// The calculated pixel movement may result in valid or invalid underflows and
	// overflows. The valid version of those boundary jumps allows players to move
	// forward to the following inner and outer buckets of the layered coordinate
	// system. The invalid version of the respective underflow and overflow
	// violations appear on the very edges of our coordinate system. Once we
	// detect such a violation we return the overflow bytes for the lower case
	// letters 't', 'r', 'b' and 'l'.  Those boundary violation bytes represent
	// the underflows and overflows towards the top, right, bottom and left
	// respectively.

	switch spc[0] {
	case 0x01:
		{
			tx2 += ds
			ty2 += dc
		}

		if tx2 > Max {
			{
				tx0, tx1 = incByt(tx0, tx1)
			}

			if tx0 > Max {
				// Overflow to the right, beyond the allowed positive x-axis boundary.
				return Bucket{}, Pixel{}, byte('r')
			}

			{
				tx2 -= Siz
			}
		}

		if ty2 > Max {
			{
				ty0, ty1 = incByt(ty0, ty1)
			}

			if ty0 > Max {
				// Overflow to the top, beyond the allowed positive y-axis boundary.
				return Bucket{}, Pixel{}, byte('t')
			}

			{
				ty2 -= Siz
			}
		}
	case 0x02:
		{
			tx2 += dc
			ty2 -= ds
		}

		if tx2 > Max {
			{
				tx0, tx1 = incByt(tx0, tx1)
			}

			if tx0 > Max {
				// Overflow to the right, beyond the allowed positive x-axis boundary.
				return Bucket{}, Pixel{}, byte('r')
			}

			{
				tx2 -= Siz
			}
		}

		if ty2 < Min {
			{
				ty0, ty1 = decByt(ty0, ty1)
			}

			if ty0 < Min {
				// Overflow to the bottom, beyond the allowed negative y-axis boundary.
				return Bucket{}, Pixel{}, byte('b')
			}

			{
				ty2 += Siz
			}
		}
	case 0x03:
		{
			tx2 -= ds
			ty2 -= dc
		}

		if tx2 < Min {
			{
				tx0, tx1 = decByt(tx0, tx1)
			}

			if tx0 < Min {
				// Underflow to the left, beyond the allowed negative x-axis boundary.
				return Bucket{}, Pixel{}, byte('l')
			}

			{
				tx2 += Siz
			}
		}

		if ty2 < Min {
			{
				ty0, ty1 = decByt(ty0, ty1)
			}

			if ty0 < Min {
				// Overflow to the bottom, beyond the allowed negative y-axis boundary.
				return Bucket{}, Pixel{}, byte('b')
			}

			{
				ty2 += Siz
			}
		}
	case 0x04:
		{
			tx2 -= dc
			ty2 += ds
		}

		if tx2 < Min {
			{
				tx0, tx1 = decByt(tx0, tx1)
			}

			if tx0 < Min {
				// Underflow to the left, beyond the allowed negative x-axis boundary.
				return Bucket{}, Pixel{}, byte('l')
			}

			{
				tx2 += Siz
			}
		}

		if ty2 > Max {
			{
				ty0, ty1 = incByt(ty0, ty1)
			}

			if ty0 > Max {
				// Overflow to the top, beyond the allowed positive y-axis boundary.
				return Bucket{}, Pixel{}, byte('t')
			}

			{
				ty2 -= Siz
			}
		}
	}

	// We return the updated position, given the current position, the direction
	// of movement, and the velocity at which a player moves during a standard
	// frame duration. The returned pixel location is rounded to simple floating
	// point precision, which means that we are only returning the nearest full
	// pixel changes of movement with mirror consistency. Code executing at this
	// point does not represent a boundary violation of our layered coordinate
	// system. Therefore we return the empty byte as third argument.

	return Bucket{tx0, ty0, tx1, ty1}, Pixel{tx2, ty2}, 0x00
}
