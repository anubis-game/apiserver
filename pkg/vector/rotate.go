package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// rotate moves this Vector along the direction of the given target coordinate,
// without changing the underlying amount of nodes.
func (v *Vector) rotate(hea matrix.Coordinate, hid byte) matrix.Coordinate {
	// Vector.rotate() can prevent extra allocations if both the head and the tail
	// nodes have to be rotated together. If either of the head or the tail node
	// is able to move further by means of accumulating a hidden node, then we
	// have to modify the head and and the tail nodes separately. This then incurs
	// extra memory allocation and garbage collection pressure.

	if v.tai.nxt.hid > 0 || v.hea.hid < v.mhn {
		v.expand(hea, hid)
		return v.shrink(hid)
	}

	// TODO:test write a unit test to verify that the head and the tail can be
	// replaced simultaneously.

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
