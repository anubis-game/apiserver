package vector

import (
	"github.com/anubis-game/apiserver/pkg/object"
)

func (v *Vector) Header() object.Object {
	return v.obj[0]
}
