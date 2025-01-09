package schema

import (
	"bytes"
)

func Decode(byt []byte) [][]byte {
	return bytes.Split(byt, Comma)
}
