package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
)

// Screen returns the partition coordinates of this Vector's current view. The
// representation of what a player sees must be maintained efficiently. This set
// of partitions changes at the outer sides of the screen while the amount of
// visible partitions remains constant. The amount of partitions may also change
// if the player grows to see farther. The input for changing the outer sides is
// the change in head partition. The input for adding an additional outer layer
// is the result of a player size dependent calculation.
func (v *Vector) Screen() []matrix.Partition {
	return v.scr
}

func (v *Vector) newScr() {
	prt := v.hea.crd.Prt()
	fos := v.crx.Fos
	pxl := fos * matrix.Prt
	ind := 0

	top := prt.Y + pxl
	rig := prt.X + pxl
	bot := prt.Y - pxl
	lef := prt.X - pxl

	{
		n := (fos * 2) + 1
		v.scr = make([]matrix.Partition, n*n)
	}

	for y := top; y >= bot; y -= matrix.Prt {
		for x := lef; x <= rig; x += matrix.Prt {
			v.scr[ind].X = x
			v.scr[ind].Y = y
			ind++
		}
	}
}
