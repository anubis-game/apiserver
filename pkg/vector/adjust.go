package vector

import (
	"math"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/object"
)

const (
	Ai float64 = 30 // Initial angle byte
	Al float64 = 1  // Largest angle byte
	Li float64 = 10 // Initial number of body parts
	// TODO:game we probably want to reduce max num to 500
	Ll float64 = 5_000 // Largest number of body parts
	// TODO:game if we half the partition size we can double the view expansion quality
	Pi float64 = 2      // Initial range of sight in partitions
	Pl float64 = 8      // Largest range of sight in partitions
	Ri float64 = 10     // Initial body part radius in pixels
	Rl float64 = 255    // Largest body part radius in pixels
	Si float64 = 50     // Initial player size in points
	Sl float64 = 50_000 // Largest player size in points
)

const (
	// Nrm is a player's normal speed.
	Nrm byte = 0x1
	// Rcn is a player's racing speed.
	Rcn byte = 0x4
)

const (
	// Frm is the standard frame duration in milliseconds travelled at a time.
	// Note that this constant is the basis for a lot of assumptions made all
	// across the game engine. Changing this value, especially decreasing it, may
	// break a number of runtime expectations.
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

// TODO:game Vector.Adjust must call Vector.Smooth() every second (once in 25 Adjust() calls)

func (v *Vector) Adjust(del int, qdr byte, agl byte, rac byte) (object.Object, object.Object, object.Object) {
	crx := v.crx.Get()

	// Increment or decrement size based on the given delta.

	{
		crx.Siz += del
	}

	var siz float64
	{
		siz = float64(crx.Siz)
	}

	// Update the player's body part radius, range of motion and range of sight.

	{
		crx.Alr = angle(siz)
		crx.Als = crx.Alr * 2
		crx.Prt = sight(siz)
		crx.Rad = radius(siz)
	}

	// Use normal speed or allow a player to race. The default is always the hard
	// coded standard speed. Only if a player provides the correct encoding for
	// racing speed, only then it will be granted. For us it is important to not
	// allow anyone to select arbitrary speed values.

	dis := Dis
	lim := crx.Als
	if rac == Rcn {
		dis = Ris
		lim = crx.Alr
	}

	// Find the allowed range of motion, given the player's current orientation
	// and desired direction to move next.

	{
		v.mot.Qdr, v.mot.Agl = trgAgl(v.mot.Qdr, v.mot.Agl, qdr, agl, lim)
	}

	// The reconciled range of motion enables us now to define the next target
	// coordinates. The updated player size will indicate if the vector has to
	// shrink, expand or only rotate towards the target coordinates.

	var len int
	{
		len = length(siz)
	}

	var hea object.Object
	{
		hea = v.Target(qdr, agl, dis)
	}

	// Adjust the player's range of sight before we actually shrink, expand or
	// rotate. The possibly modified range of sight defines the list of partitions
	// that we eventually move towards in either direction. So we have to compute
	// the partition boundaries of the player's screen before we advance the
	// vector header within the coordinate system.

	var prt object.Object
	{
		prt = hea.Prt()
	}

	var vpb int
	{
		vpb = crx.Prt * matrix.Prt
	}

	// We reset the existing *Screen pointer so that we do not allocate all the
	// time. This safes about 17 ns/op compared to creating a new struct pointer.

	{
		v.scr.Top = prt.Y + vpb
		v.scr.Rig = prt.X + vpb
		v.scr.Bot = prt.Y - vpb
		v.scr.Lef = prt.X - vpb
		v.scr.Prt = nil
	}

	if len > v.len {
		v.Expand(hea)
	}

	// TODO:test ensure tl1 and tl2 are constructed properly per call to Vector.Adjust()

	var tl1 object.Object
	if len == v.len {
		tl1 = v.Rotate(hea)
	}

	var tl2 object.Object
	if len < v.len {
		tl2 = v.Shrink()
	}

	// Track the latest character settings according to the reconciliation between
	// our system rules and the desired state.

	{
		v.crx.Set(crx)
	}

	return hea, tl1, tl2
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
	len := math.Max(Li, Li*math.Pow(siz/Si, 0.9))
	return int(math.Min(Ll, math.Ceil(len)))
}

// See https://www.desmos.com/calculator/dokimlkswz for probably outdated radius
// calculations.
func radius(siz float64) byte {
	rad := math.Max(Ri, Ri*math.Pow(siz/Si, 0.47))
	return byte(math.Min(Rl, math.Ceil(rad)))
}

// See https://www.desmos.com/calculator/vpiapfimb1 for probably outdated sight
// calculations.
func sight(siz float64) int {
	rad := math.Max(Pi, Pi*math.Pow(siz/Si, 0.2))
	return int(math.Min(Pl, math.Round(rad)))
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
