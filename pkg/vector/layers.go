package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// TODO:test make sure Header does not panic

// Layers returns a symmetric set of partition coordinates around the current
// head node of this Vector, based on the provided factor of sight and partition
// length. E.g. fos=len, 1=9, 2=25, 3=47 etc.
func (v *Vector) Layers(fos int, prt int) []matrix.Partition {
	var chp matrix.Partition
	if prt == matrix.Pt8 {
		chp = v.hea.crd.Pt8()
	} else {
		chp = v.hea.crd.Pt1()
	}

	pxl := fos * prt

	top := chp.Y + pxl
	rig := chp.X + pxl
	bot := chp.Y - pxl
	lef := chp.X - pxl

	ind := 0
	row := (fos * 2) + 1
	scr := make([]matrix.Partition, row*row)

	for y := top; y >= bot; y -= prt {
		for x := lef; x <= rig; x += prt {
			scr[ind].X = x
			scr[ind].Y = y
			ind++
		}
	}

	return scr
}
