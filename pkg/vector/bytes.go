package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
)

func (v *Vector) Bytes() []byte {
	// TODO
	return nil
}

func FromBytes(byt []byte) *Vector {
	if len(byt)%6 != 0 {
		panic(fmt.Sprintf("invalid vector byte length: %d", len(byt)))
	}

	var obj []object.Object

	for i := 0; i < len(byt)/6; i++ {
		pos := i * 6
		obj = append(obj, object.FromBytes(byt[pos:pos+6]))
	}

	return New(Config{
		Obj: obj,
	})
}
