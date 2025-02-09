package window

import (
	"github.com/anubis-game/apiserver/pkg/object"
)

type Window struct {
	cbl object.Object
	ctr object.Object
}

// Has returns whether the given bucket resides inside the underlying Window. So
// if obj turns out to be outside of w, then Has returns false.
func (w *Window) Has(obj object.Object) bool {
	return obj.X >= w.cbl.X && obj.X <= w.ctr.X && obj.Y >= w.cbl.Y && obj.Y <= w.ctr.Y
}
