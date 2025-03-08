package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// rotate moves this Vector along the direction of the given target coordinate,
// without changing the underlying amount of nodes.
func (v *Vector) rotate(hea matrix.Coordinate) matrix.Coordinate {
	cur := v.tai   // remember tail pointer as new head pointer
	prv := cur.crd // old tail coordinates get removed

	{
		v.tai = cur.nxt // next of tail becomes new tail
		v.tai.prv = nil // remove previous of new tail
	}

	{
		cur.nxt = nil   // nil next of new head
		cur.crd = hea   // set value of new head
		cur.prv = v.hea // point back to old head
	}

	{
		v.hea.nxt = cur // old head links to new head
		v.hea = cur     // old tail becomes new head
	}

	return prv
}
