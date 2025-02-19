package vector

import (
	"github.com/anubis-game/apiserver/pkg/object"
)

func (v *Vector) Encode() []byte {
	buf := make([]byte, 8+v.len*object.Len)

	{
		copy(buf[0:2], v.uid[:])
	}

	{
		crx := v.Charax().Get()
		buf[2] = byte(crx.Rad)
		buf[3] = byte(crx.Siz)
		buf[4] = crx.Typ
	}

	{
		mot := v.Motion().Get()
		buf[5] = mot.Qdr
		buf[6] = mot.Agl
		buf[7] = mot.Vlc
	}

	cur := v.tai
	ind := 0
	for cur != nil {
		pos := 8 + (ind * object.Len)
		byt := cur.val.Byt()

		copy(buf[pos:pos+object.Len], byt[:])

		cur = cur.nxt
		ind++
	}

	return buf
}
