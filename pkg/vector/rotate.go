package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// rotate moves the vector along the direction of the given target coordinates
// without expanding the underlying amount of segments.
func (v *Vector) rotate(hea matrix.Coordinate, hid byte) matrix.Coordinate {
	// Vector.rotate() can prevent extra allocations if both the head and the tail
	// segments have to be rotated together. If either of the head or the tail
	// segment is able to move by accumulating a hidden segment, then we have to
	// expand the head and shrink the tail separately, which incurs extra memory
	// allocation and garbage collection pressure.

	if v.tai.nxt.hid > 0 || v.hea.hid < v.mhs {
		v.expand(hea, hid)
		return v.shrink(hid)
	}

	new := v.tai   // remember tail pointer as new head pointer
	old := new.crd // old tail coordinates get removed

	{
		v.tai = new.nxt // next of tail becomes new tail
		v.tai.prv = nil // remove previous of new tail
	}

	{
		new.nxt = nil   // nil next of new head
		new.crd = hea   // set value of new head
		new.prv = v.hea // point back to old head
	}

	{
		v.hea.nxt = new // old head links to new head
		v.hea = new     // old tail becomes new head
	}

	return old
}
