package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
)

func (v *Vector) Bytes() []byte {
	byt := make([]byte, v.siz*6)

	cur := v.tai
	ind := 0
	for cur != nil {
		f := ind * 6
		b := cur.val.Bucket()
		copy(byt[f:f+6], b[:])
		cur = cur.nxt
		ind++
	}

	return byt
}

func FromBytes(byt []byte) *Vector {
	if len(byt)%6 != 0 {
		panic(fmt.Sprintf("invalid vector byte length: %d", len(byt)))
	}

	var obj []object.Object

	for i := 0; i < len(byt)/6; i++ {
		f := i * 6
		obj = append(obj, object.Bucket(byt[f:f+6]).Object())
	}

	return New(Config{
		Obj: obj,
	})
}
