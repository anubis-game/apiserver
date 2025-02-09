package vector

import (
	"github.com/anubis-game/apiserver/pkg/object"
)

// Buffer returns the prepared fanout buffer for the given coordinate partition.
func (v *Vector) Buffer(prt object.Object) []byte {
	return v.buf[prt]
}
