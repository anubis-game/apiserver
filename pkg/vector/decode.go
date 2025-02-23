package vector

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
)

// decode is only used for testing Vector encoding.
func decode(byt []byte) *Vector {
	if (len(byt)-3)%object.Len != 0 {
		panic(fmt.Sprintf("invalid vector byte length: %d", len(byt)))
	}

	var obj []object.Object
	for i := range (len(byt) - 3) / object.Len {
		pos := 3 + (i * object.Len)
		obj = append(obj, object.New(byt[pos:pos+object.Len]))
	}

	return New(Config{
		Obj: obj,
		Uid: byt[1],
	})
}
