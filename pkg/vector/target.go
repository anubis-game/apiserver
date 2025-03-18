package vector

import (
	"math"
)

const (
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
	cos [256]float64
	// sin is the sine lookup table to cache all possible sine values based on any
	// given angle byte.
	sin [256]float64
)

func init() {
	for i := range 256 {
		cos[i] = math.Cos(float64(i) * qrf)
		sin[i] = math.Sin(float64(i) * qrf)
	}
}

func target(chx int, chy int, qdr byte, agl byte, dis float64) (int, int) {
	// Update both coordinates based on the current position and the desired
	// direction.

	switch qdr {
	case 0x1:
		chx += int(dis*sin[agl] + 0.5)
		chy += int(dis*cos[agl] + 0.5)
	case 0x2:
		chx += int(dis*cos[agl] + 0.5)
		chy -= int(dis*sin[agl] + 0.5)
	case 0x3:
		chx -= int(dis*sin[agl] + 0.5)
		chy -= int(dis*cos[agl] + 0.5)
	case 0x4:
		chx -= int(dis*cos[agl] + 0.5)
		chy += int(dis*sin[agl] + 0.5)
	}

	return chx, chy
}
