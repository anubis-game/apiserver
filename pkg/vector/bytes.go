package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
)

func (v *Vector) Bytes() []byte {
	byt := make([]byte, v.len*6)

	for i := range v.len {
		f := i * 6
		b := v.obj[i].Bucket()
		copy(byt[f:f+6], b[:])
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
