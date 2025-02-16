package energy

import (
	"github.com/anubis-game/apiserver/pkg/object"
)

func (e *Energy) Encode() []byte {
	var buf [8]byte

	b := e.Obj.Byt()
	copy(buf[0:object.Len], b[:])

	buf[6] = e.Siz
	buf[7] = e.Typ

	return buf[:]
}
