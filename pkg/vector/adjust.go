package vector

import (
	"math"

	"github.com/anubis-game/apiserver/pkg/object"
)

const (
	Ai float64 = 30     // Initial angle byte
	Al float64 = 1      // Largest angle byte
	Li float64 = 10     // Initial number of body parts
	Ll float64 = 5_000  // Largest number of body parts
	Pi float64 = 3      // Initial range of sight in partitions
	Pl float64 = 25     // Largest range of sight in partitions
	Ri float64 = 10     // Initial body part radius in pixels
	Rl float64 = 256    // Largest body part radius in pixels
	Si float64 = 50     // Initial player size in points
	Sl float64 = 50_000 // Largest player size in points
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

func (v *Vector) Adjust(del int, des Motion) {
	crx := v.crx.Get()
	cur := v.mot.Get()

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
	if des.Vlc == Rcn {
		dis = Ris
		lim = crx.Alr
	}

	// Find the allowed range of motion, given the player's current orientation
	// and desired direction to move next.

	{
		des.Qdr, des.Agl = trgAgl(cur.Qdr, cur.Agl, des.Qdr, des.Agl, lim)
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
		hea = v.Target(des.Qdr, des.Agl, dis)
	}

	if len < v.len {
		v.Shrink()
	}

	if len > v.len {
		v.Expand(hea)
	} else {
		v.Rotate(hea)
	}

	// Track the latest range of motion and character settings according to the
	// reconciliation between our system rules and the desired state.

	{
		v.crx.Set(crx)
		v.mot.Set(des)
	}
}

// See https://www.desmos.com/calculator/kni7qb0o9y for full turn diameters.
func angle(siz float64) byte {
	agl := math.Min(1, 1-math.Pow(siz/Sl, 0.4)) * Ai
	return byte(math.Max(Al, math.Ceil(agl)))
}

// See https://www.desmos.com/calculator/7bntdys5na for length calculations.
func length(siz float64) int {
	len := math.Max(Li, Li*math.Pow(siz/Si, 0.9))
	return int(math.Min(Ll, math.Ceil(len)))
}

// See https://www.desmos.com/calculator/dokimlkswz for radius calculations.
func radius(siz float64) int {
	rad := math.Max(Ri, Ri*math.Pow(siz/Si, 0.47))
	return int(math.Min(Rl, math.Ceil(rad)))
}

// See https://www.desmos.com/calculator/vpiapfimb1 for sight calculations.
func sight(siz float64) int {
	prt := math.Max(Pi, Pi+math.Pow(siz/Si, 0.46)-1)
	return int(math.Min(Pl, math.Ceil(prt)))
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
