package energy

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/matrix"
)

// decode is only used for testing Energy encoding.
func decode(byt []byte) Config {
	if len(byt) != 8 {
		panic(fmt.Sprintf("expected 8 energy bytes, got %d", len(byt)))
	}

	return Config{
		Crd: matrix.NewCoordinate(byt),
		Siz: byt[6],
		Typ: byt[7],
	}
}
