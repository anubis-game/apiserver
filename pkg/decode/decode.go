package decode

import (
	"bytes"
)

const (
	Sep = byte(',')
)

func Decode(byt []byte) [][]byte {
	return bytes.Split(byt, []byte{Sep})
}
