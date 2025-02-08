package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Rotate moves the vector along the direction of the given target object
// without expanding the underlying amount of segments. After calling Rotate,
// the underlying vector has the same amount of objects as it had before.
func (v *Vector) Rotate(hea object.Object) {
	tai := v.tai   // remember tail as new head
	old := tai.val // the old tail gets cleaned up

	{
		v.tai = tai.nxt // next of tail becomes new tail
	}

	{
		tai.nxt = nil // nil next of new head
		tai.val = hea // set value of new head
	}

	{
		v.hea.nxt = tai // old head links to new head
		v.hea = tai     // old tail becomes new head
	}

	{
		v.expand(hea)
		v.shrink(old)
	}
}
