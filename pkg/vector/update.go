package vector

import (
	"math"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

const (
	Ai float64 = 30     // Initial angle byte
	Al float64 = 1      // Largest angle byte
	Li float64 = 2      // Initial number of body parts
	Ll float64 = 2_500  // Largest number of body parts
	Ri float64 = 10     // Initial body part radius in pixels
	Rl float64 = 225    // Largest body part radius in pixels
	Si float64 = 10     // Initial player size in points
	Sl float64 = 50_000 // Largest player size in points
)

const (
	// Nrm is a player's normal speed.
	Nrm byte = 0x1
	// Rcn is a player's racing speed.
	Rcn byte = 0x2
)

const (
	// Frm is the standard frame duration in milliseconds travelled at a time.
	// Note that this constant is the basis for a lot of assumptions made all
	// across the game engine. Changing this value, especially decreasing it, may
	// break a number of runtime expectations.
	Frm = 25
)

const (
	// nrm is the standard distance travelled in pixels per standard frame. The
	// distance travelled per frame at normal speed is 5 pixels.
	nrm float64 = 5
	// rcn is the increased distance travelled in pixels per standard frame. The
	// distance travelled per frame at racing speed is 10 pixels.
	rcn float64 = float64(Rcn) * nrm
)

// TODO:game Vector.Adjust must call Vector.Smooth() every second (once in 25 Adjust() calls)

func (v *Vector) Update(del int, qdr byte, agl byte, rac byte) {
	// Increment or decrement size based on the given delta.

	{
		v.crx.Siz += del
	}

	var siz float64
	{
		siz = float64(v.crx.Siz)
	}

	// Get the player's desired factor of sight, given the new player size. This
	// factor here will be multiplied by the length of our matrix partition like
	// shown below.
	//
	//     ((fos * 2) + 1) * matrix.Prt
	//

	var fos int
	{
		fos = sight(siz)
	}

	// We have to check whether the player screen has to be expanded or shrunk
	// based on the current and desired factor of sight. If those two values
	// differ then we have to re-allocate the Vector's screen partitions. This
	// re-allocation may mean to grow or shrink the amount of partition
	// coordinates visible to the player.

	var rsz bool
	if v.crx.Fos != fos {
		rsz = true
	}

	// Update the rest of the character settings.

	{
		v.crx.Alr = angle(siz)
		v.crx.Aln = v.crx.Alr * 2
		v.crx.Fos = fos
		v.crx.Rad = radius(siz)
	}

	// Use normal speed or allow a player to go fast in racing mode. The default
	// is always the hard coded normal speed. Only if a player provides the
	// correct encoding for racing speed, only then will the faster movement be
	// granted, because we cannot allow anyone to select arbitrary speed values.

	dis := nrm
	hid := Nrm
	lim := v.crx.Aln
	if rac == Rcn {
		dis = rcn
		hid = Rcn
		lim = v.crx.Alr
	}

	// Find the allowed range of motion, given the player's current orientation
	// and desired direction of travel. Smaller players have a smaller turning
	// radius. Larger player require more time to turn around.

	{
		v.mot.Qdr, v.mot.Agl = trgAgl(v.mot.Qdr, v.mot.Agl, qdr, agl, lim)
	}

	// The reconciled range of motion enables us now to define the next target
	// coordinates. The updated player size will indicate if the Vector has to
	// shrink, expand or only rotate towards the new target coordinates.

	var hea matrix.Coordinate
	{
		hea = v.target(v.mot.Qdr, v.mot.Agl, dis)
	}

	var len int
	{
		len = length(siz)
	}

	// Reset the occupied coordinate diff before we add to it below. This enables
	// us to only track the change of coordinate movements that this Vector
	// applied during this update cycle.

	{
		v.ocd = map[matrix.Partition][]matrix.Coordinate{}
	}

	// Shrink if we are too long. Note that we shrink first so that we do not
	// create observational side effects for Vector.expand() and Vector.rotate().

	if len < v.len {
		v.len--
		old := v.shrink(hid)
		v.occRem(old)
	}

	// At this point a tail node may or may not have been already removed. We can
	// now either expand if we are supposed to grow larger, or simply rotate
	// forward.

	if len > v.len {
		v.len++
		v.expand(hea, hid)
		v.occAdd(rsz)
	} else {
		old := v.rotate(hea, hid)
		v.occAdd(rsz)
		v.occRem(old)
	}

	if rsz {
		v.newScr()
	}
}

// See https://www.desmos.com/calculator/kni7qb0o9y for probably outdated full
// turn diameters.
func angle(siz float64) byte {
	agl := math.Min(1, 1-math.Pow(siz/Sl, 0.4)) * Ai
	return byte(math.Max(Al, math.Ceil(agl)))
}

// See https://www.desmos.com/calculator/7bntdys5na for probably outdated length
// calculations.
func length(siz float64) int {
	len := math.Max(Li, Li*math.Pow(siz/Si, 0.8372))
	return int(math.Min(Ll, math.Ceil(len)))
}

// See https://www.desmos.com/calculator/dweegd1sxr for the radius calculations.
func radius(siz float64) byte {
	rad := math.Max(Ri, math.Pow(math.Log10(Ri*(siz/Si)), 3.5))
	return byte(math.Min(Rl, math.Ceil(rad)))
}

func sight(siz float64) int {
	switch s := siz; {
	case s < 250: //    between 0 and 249
		return 2
	case s < 1_000: //  between 250 and 999
		return 3
	case s < 5_000: //  between 1,000 and 4,999
		return 4
	case s < 10_000: // between 5,000 and 9,999
		return 5
	case s < 25_000: // between 10,000 and 24,999
		return 6
	default: //         above 25,000
		return 7
	}
}

func trgAgl(cqd byte, cag byte, dqd byte, dag byte, lim byte) (byte, byte) {
	// Convert the given inputs to absolute angles on the basis of a full
	// 1024-step circle.

	acr := (int(cqd)-1)*256 + int(cag)
	ads := (int(dqd)-1)*256 + int(dag)

	// Compute the desired displacement in either direction.

	dif := ads - acr

	// Resolve the approriate direction and apply the maximum allowed range of
	// motion.

	var anw int
	if dif < -int(lim) || dif > 512 {
		anw = (acr - int(lim) + 1024) % 1024
	} else if dif > int(lim) || dif < -512 {
		anw = (acr + int(lim) + 1024) % 1024
	} else {
		anw = (acr + dif + 1024) % 1024
	}

	// Translate the absolute result into the final quadrant and angle.

	return byte(anw/256 + 1), byte(anw % 256)
}
