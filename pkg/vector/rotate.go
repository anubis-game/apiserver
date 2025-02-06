package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Rotate moves the vector along the direction of the given target object
// without expanding the underlying amount of segments. After calling Rotate,
// the underlying vector has the same amount of objects as it had before.
func (v *Vector) Rotate(trg object.Object) {
	{
		v.lis.PushFront(trg)
	}

	rem := v.lis.Remove(v.lis.Back()).(object.Object)

	{
		v.win.Inc(trg)
		v.win.Dec(v.lis.Back().Value.(object.Object), rem)
	}
}
