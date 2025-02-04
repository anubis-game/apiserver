package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Rotate moves the vector along the direction of the given target object
// without expanding the underlying amount of segments. After calling Rotate,
// the underlying vector has the same amount of objects as it had before.
func (v *Vector) Rotate(trg object.Object) {
	ind := v.ind - 1
	tai := v.obj[ind]

	{
		copy(v.obj[1:], v.obj[:ind]) // shift without tail
		v.obj[0] = trg               // target becomes head
	}

	{
		v.win.Inc(trg)
		v.win.Dec(v.obj[ind], tai)
	}
}
