package decode

import (
	"bytes"
)

var (
	Sep = []byte{','}
)

func Decode(byt []byte) [][]byte {
	return bytes.Split(byt, Sep)
}
