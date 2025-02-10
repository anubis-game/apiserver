package vector

import (
	"math"

	"github.com/anubis-game/apiserver/pkg/object"
)

const (
	// HD is the index for the head pixel of the vector chain.
	HD int = 0
)

const (
	// Frm is the standard frame duration in milliseconds travelled at a time.
	Frm = 25
	// Dis is the standard distance travelled in pixels per standard frame. The
	// amount of pixels travelled here per millisecond is 0.2, which represents a
	// velocity of 1.
	Dis float64 = Frm * 0.2
	// Ris is the increased distance travelled in pixels per standard frame. The
	// amount of pixels travelled here per millisecond is 0.8, which represents a
	// velocity of 4.
	Ris float64 = Frm * 0.8
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
	for i := 0; i < 256; i++ {
		cos[i] = math.Cos(float64(i) * qrf)
		sin[i] = math.Sin(float64(i) * qrf)
	}
}

func (v *Vector) Target(mot Motion) object.Object {
	prv := v.mot.Get()
	tpx := v.hea.val

	// Use normal speed or allow a player to race. The default is always the hard
	// coded standard speed. Only if a player provides the correct encoding for
	// racing speed, only then it will be granted. For us it is important to not
	// allow anyone to select arbitrary speed values.

	// TODO configure motion range limit dynamically
	dis := Dis
	lim := byte(100)
	if mot.Vlc == 0x4 {
		dis = Ris
		lim = byte(50)
	} else {
		mot.Vlc = 0x1
	}

	{
		mot.Qdr, mot.Agl = trgAgl(prv.Qdr, prv.Agl, mot.Qdr, mot.Agl, lim)
	}

	// Update both coordinates based on the current position and the desired
	// direction.

	switch mot.Qdr {
	case 0x1:
		tpx.X += int(dis*sin[mot.Agl] + 0.5)
		tpx.Y += int(dis*cos[mot.Agl] + 0.5)
	case 0x2:
		tpx.X += int(dis*cos[mot.Agl] + 0.5)
		tpx.Y -= int(dis*sin[mot.Agl] + 0.5)
	case 0x3:
		tpx.X -= int(dis*sin[mot.Agl] + 0.5)
		tpx.Y -= int(dis*cos[mot.Agl] + 0.5)
	case 0x4:
		tpx.X -= int(dis*cos[mot.Agl] + 0.5)
		tpx.Y += int(dis*sin[mot.Agl] + 0.5)
	}

	// Track the latest range of motion according to the reconciliation between
	// our system rules and the desired state.

	{
		v.mot.Set(mot)
	}

	return tpx
}

func trgAgl(pqd byte, pag byte, nqd byte, nag byte, lim byte) (byte, byte) {
	// Convert the given inputs to absolute angles on the basis of a full
	// 1024-step circle.

	apv := (int(pqd)-1)*256 + int(pag)
	anx := (int(nqd)-1)*256 + int(nag)

	// Compute the desired displacement in either direction.

	dif := anx - apv

	// Resolve the approriate direction and apply the maximum allowed range of
	// motion.

	var anw int
	if dif < -int(lim) || dif > 512 {
		anw = (apv - int(lim) + 1024) % 1024
	} else if dif > int(lim) || dif < -512 {
		anw = (apv + int(lim) + 1024) % 1024
	} else {
		anw = (apv + dif + 1024) % 1024
	}

	// Translate the absolute result into the final quadrant and angle.

	return byte(anw/256 + 1), byte(anw % 256)
}
