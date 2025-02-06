package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Rotate moves the vector along the direction of the given target object
// without expanding the underlying amount of segments. After calling Rotate,
// the underlying vector has the same amount of objects as it had before.
func (v *Vector) Rotate(val object.Object) {
	tai := v.tai // remember tail as new head
	rem := tai.val

	{
		v.tai = tai.nxt // next of tail becomes new tail
	}

	{
		tai.nxt = nil // nil next of new head
		tai.val = val // set value of new head
	}

	{
		v.hea.nxt = tai // old head links to new head
		v.hea = tai     // old tail becomes new head
	}

	// TODO why is this so expensive and do we even need the window magic here?
	{
		v.win.Inc(val)
		v.win.Dec(v.tai.val, rem)
	}
}
