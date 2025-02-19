package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
)

// decode is only used for testing Vector encoding.
func decode(byt []byte) *Vector {
	if (len(byt)-8)%object.Len != 0 {
		panic(fmt.Sprintf("invalid vector byte length: %d", len(byt)))
	}

	var obj []object.Object
	for i := 0; i < (len(byt)-8)/object.Len; i++ {
		pos := 8 + (i * object.Len)
		obj = append(obj, object.New(byt[pos:pos+object.Len]))
	}

	var mot Motion
	{
		mot = Motion{
			Qdr: byt[5],
			Agl: byt[6],
			Vlc: byt[7],
		}
	}

	var vec *Vector
	{
		vec = New(Config{
			Mot: mot,
			Obj: obj,
			Uid: [2]byte{byt[0], byt[1]},
		})
	}

	{
		vec.Charax().Set(Charax{
			Rad: int(byt[2]),
			Siz: int(byt[3]),
			Typ: byt[4],
		})
	}

	return vec
}
