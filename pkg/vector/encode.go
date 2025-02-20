package vector

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/schema"
)

func (v *Vector) Encode() []byte {
	var buf []byte
	{
		buf = make([]byte, 4+v.len*object.Len)
	}

	{
		buf[0] = byte(schema.Body)
		copy(buf[1:3], v.uid[:])
		buf[3] = byte(v.len)
	}

	cur := v.tai
	ind := 0
	for cur != nil {
		pos := 4 + (ind * object.Len)
		byt := cur.val.Byt()

		copy(buf[pos:pos+object.Len], byt[:])

		cur = cur.nxt
		ind++
	}

	return buf
}
