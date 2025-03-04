package energy

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

// decode is only used for testing Energy encoding.
func decode(byt []byte) *Energy {
	if len(byt) != 9 {
		panic(fmt.Sprintf("expected 9 energy bytes, got %d", len(byt)))
	}

	return New(Config{
		Obj: matrix.NewCoordinate(byt[1 : 1+matrix.CoordinateBytes]),
		Siz: byt[7],
		Typ: byt[8],
	})
}
