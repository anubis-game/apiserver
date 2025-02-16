package vector

import (
	"github.com/anubis-game/apiserver/pkg/object"
)

func (v *Vector) Encode() []byte {
	buf := make([]byte, v.len*object.Len)

	cur := v.tai
	ind := 0
	for cur != nil {
		pos := ind * object.Len
		byt := cur.val.Byt()

		copy(buf[pos:pos+object.Len], byt[:])

		cur = cur.nxt
		ind++
	}

	return buf
}
