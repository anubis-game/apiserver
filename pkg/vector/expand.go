package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// expand moves the vector along the direction of the given target object and
// expands the underlying segments. After calling Expand, the underlying vector
// has 1 more object, which is the added target head.
func (v *Vector) expand(hea matrix.Coordinate, hid byte) {
	if v.hea.hid < v.mhs {
		v.hea.hid += int8(hid)
		v.hea.crd = hea
	} else {
		lin := &Linker{
			crd: hea,
			prv: v.hea,
		}

		{
			v.hea.nxt = lin
			v.hea = lin
		}
	}
}
