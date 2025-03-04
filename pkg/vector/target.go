package vector

import (
	"math"

	"github.com/anubis-game/apiserver/pkg/matrix"
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

func (v *Vector) target(qdr byte, agl byte, dis float64) matrix.Coordinate {
	trg := v.hea.crd

	// Update both coordinates based on the current position and the desired
	// direction.

	switch qdr {
	case 0x1:
		trg.X += int(dis*sin[agl] + 0.5)
		trg.Y += int(dis*cos[agl] + 0.5)
	case 0x2:
		trg.X += int(dis*cos[agl] + 0.5)
		trg.Y -= int(dis*sin[agl] + 0.5)
	case 0x3:
		trg.X -= int(dis*sin[agl] + 0.5)
		trg.Y -= int(dis*cos[agl] + 0.5)
	case 0x4:
		trg.X -= int(dis*cos[agl] + 0.5)
		trg.Y += int(dis*sin[agl] + 0.5)
	}

	return trg
}
