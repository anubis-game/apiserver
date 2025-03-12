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

// TODO:game Vector.Update() must call Vector.Smooth() every second or so (once
// in 25 calls).

func (v *Vector) Update(del int, qdr byte, agl byte, rac byte) {
	// Increment or decrement size based on the given delta.

	{
		v.crx.Siz += del
	}

	var siz float64
	{
		siz = float64(v.crx.Siz)
	}

	// Update the player's character settings.

	{
		v.crx.Alr = angle(siz)
		v.crx.Aln = v.crx.Alr * 2
		v.crx.Fos = sight(siz)
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
	// coordinates. We also remember the current tail so that we can track the
	// changes properly, that we applied during this update cycle.

	var hea matrix.Coordinate
	{
		hea = v.target(v.mot.Qdr, v.mot.Agl, dis)
	}

	var ta1 matrix.Coordinate
	var ta2 matrix.Coordinate
	var ta3 matrix.Coordinate
	{
		ta1 = v.tai.crd
	}

	// The updated player size will indicate whether the Vector has to shrink,
	// expand or only rotate towards the new target coordinates.

	var nod int
	{
		nod = length(siz)
	}

	// Shrink if we are too long. Note that we shrink first so that we do not
	// create observational side effects for Vector.expand() and Vector.rotate().

	if nod < v.len {
		v.len--
		ta2 = v.shrink(hid)
		v.occRem(ta2)
	}

	// At this point a tail node may or may not have been already removed. We can
	// now either expand if we are supposed to grow larger, or simply rotate
	// forward.
	// Vector.rotate() can prevent extra allocations if both the head and the tail
	// nodes have to be rotated together. If either of the head or the tail node
	// is able to move further by means of accumulating a hidden node, then we
	// have to modify the head and and the tail nodes separately. This then incurs
	// extra memory allocation and garbage collection pressure.

	if nod > v.len {
		v.len++
		v.expand(hea, hid)
		v.occAdd()
	} else if v.tai.nxt.hid > 0 || v.hea.hid < v.mhn {
		v.expand(hea, hid)
		ta3 = v.shrink(hid)
		v.occAdd()
		v.occRem(ta3)
	} else {
		ta3 = v.rotate(hea)
		v.occAdd() // TODO:test we have no coverage on screen and occupied boundaries
		v.occRem(ta3)
	}

	// Reset the occupied coordinate diff with the change of coordinate movements
	// that this Vector applied during this update cycle. There is always a new
	// head to track. There may be a new tail if we rotated or shrunk. There may
	// be an old tail if we just rotated. And there may be another old tail if we
	// rotated and shrunk.

	if v.ocd.Rem != nil {
		v.ocd.Rem = nil
	}

	{
		v.ocd.Hea = v.hea.crd
	}

	if !ta1.Eql(v.tai.crd) {
		v.ocd.Tai = v.tai.crd
	} else {
		v.ocd.Tai = matrix.Coordinate{}
	}

	// If we shrink and rotate within the same update cycle, then we remove two
	// tail nodes. If we then care about the order of changed tail nodes to
	// consume using Vector.Change(), then we have to add ta3 to the delta
	// directly after the new tail node. If ta2 exists, add that after ta3.

	if !ta3.Zer() {
		v.ocd.Rem = append(v.ocd.Rem, ta3)
	}

	if !ta2.Zer() {
		v.ocd.Rem = append(v.ocd.Rem, ta2)
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
