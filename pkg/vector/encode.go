package vector

import (
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/schema"
)

func (v *Vector) Encode() []byte {
	var buf []byte
	{
		buf = make([]byte, 3+v.len*matrix.CoordinateBytes)
	}

	{
		buf[0] = byte(schema.Body)
		buf[2] = byte(v.len)
	}

	cur := v.tai
	ind := 0
	for cur != nil {
		pos := 3 + (ind * matrix.CoordinateBytes)
		byt := cur.crd.Byt()

		copy(buf[pos:pos+matrix.CoordinateBytes], byt[:])

		cur = cur.nxt
		ind++
	}

	return buf
}
