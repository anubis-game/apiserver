package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// expand moves this Vector along the direction of the given target coordinate
// and grows the number of nodes if necessary.
func (v *Vector) expand(hea matrix.Coordinate, hid byte) {
	if v.hea.hid < v.mhn {
		v.hea.hid += int8(hid)
		v.hea.crd = hea

		return
	}

	lnk := &Linker{
		crd: hea,
		prv: v.hea,
	}

	{
		v.hea.nxt = lnk
		v.hea = lnk
	}
}
