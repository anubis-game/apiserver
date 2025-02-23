package unique

import (
	"fmt"
	"math"
)

type Number interface {
	uint8 | uint16 | int8 | int16
}

func length[T Number]() int {
	var z T

	switch t := any(z).(type) {
	case uint8:
		return math.MaxUint8
	case uint16:
		return math.MaxUint16
	case int8:
		return math.MaxInt8
	case int16:
		return math.MaxInt16
	default:
		panic(fmt.Sprintf("invalid type %T", t))
	}
}
