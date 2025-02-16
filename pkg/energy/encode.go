package energy

import (
	"github.com/anubis-game/apiserver/pkg/object"
)

func (e *Energy) Encode() []byte {
	byt := make([]byte, 8)

	b := e.Obj.Byt()
	copy(byt[:object.Len], b[:])

	byt[6] = e.Siz
	byt[7] = e.Typ

	return byt
}
