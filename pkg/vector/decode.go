package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Decode(byt []byte) *Vector {
	if len(byt)%object.Len != 0 {
		panic(fmt.Sprintf("invalid vector byte length: %d", len(byt)))
	}

	var obj []object.Object

	for i := 0; i < len(byt)/object.Len; i++ {
		pos := i * object.Len
		obj = append(obj, object.New(byt[pos:pos+object.Len]))
	}

	return New(Config{
		Obj: obj,
	})
}
