package energy

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/object"
)

func Decode(byt []byte) *Energy {
	if len(byt) != 8 {
		panic(fmt.Sprintf("expected 8 energy bytes, got %d", len(byt)))
	}

	return &Energy{
		Obj: object.New(byt[:object.Len]),
		Siz: byt[6],
		Typ: byt[7],
	}
}
