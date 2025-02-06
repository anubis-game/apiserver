package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Expand moves the vector along the direction of the given target object and
// expands the underlying segments. After calling Expand, the underlying vector
// has 1 more object, which is the added target head.
func (v *Vector) Expand(hea object.Object) {
	lin := &Linker{val: hea}

	v.hea.nxt = lin
	v.hea = lin
	v.siz++

	v.win.Inc(hea)
}
