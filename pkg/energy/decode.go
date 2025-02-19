package energy

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
)

// decode is only used for testing Energy encoding.
func decode(byt []byte) *Energy {
	if len(byt) != 9 {
		panic(fmt.Sprintf("expected 9 energy bytes, got %d", len(byt)))
	}

	return New(Config{
		Obj: object.New(byt[1 : 1+object.Len]),
		Siz: byt[7],
		Typ: byt[8],
	})
}
