package vector

import (
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/schema"
)

func (v *Vector) Encode() []byte {
	var buf []byte
	{
		buf = make([]byte, 3+v.len*object.Len)
	}

	{
		buf[0] = byte(schema.Body)
		buf[1] = v.uid
		buf[2] = byte(v.len)
	}

	cur := v.tai
	ind := 0
	for cur != nil {
		pos := 3 + (ind * object.Len)
		byt := cur.val.Byt()

		copy(buf[pos:pos+object.Len], byt[:])

		cur = cur.nxt
		ind++
	}

	return buf
}
