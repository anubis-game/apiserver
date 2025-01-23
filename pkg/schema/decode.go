package schema

import (
	"bytes"
)

var (
	Comma = []byte(",")
)

func Decode(byt []byte) [][]byte {
	return bytes.Split(byt, Comma)
}
