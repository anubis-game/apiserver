package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Rotate moves the vector along the direction of the given target object
// without expanding the underlying amount of segments. After calling Rotate,
// the underlying vector has the same amount of objects as it had before.
func (v *Vector) Rotate(hea object.Object) object.Object {
	new := v.tai   // remember tail as new head
	old := new.val // old tail gets cleaned up

	{
		v.tai = new.nxt // next of tail becomes new tail
	}

	{
		new.nxt = nil // nil next of new head
		new.val = hea // set value of new head
	}

	{
		v.hea.nxt = new // old head links to new head
		v.hea = new     // old tail becomes new head
	}

	{
		v.expand(hea)
		v.shrink(old)
	}

	return old
}
