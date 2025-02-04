package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Expand moves the vector along the direction of the given target object and
// expands the underlying segments. After calling Expand, the underlying vector
// has 1 more object, which is the added target head.
func (v *Vector) Expand(trg object.Object) {
	{
		copy(v.obj[1:], v.obj[:v.ind]) // shift all segments
		v.obj[0] = trg                 // target becomes head
		v.ind++                        // increment real length
	}

	{
		v.win.Inc(trg)
	}
}
