package address

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
)

var (
	zeroAddress common.Address
)

var (
	zeroBytes = []byte{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00,
	}
)

func Empty(a common.Address) bool {
	return a == zeroAddress
}

func EmptyBytes(b []byte) bool {
	return bytes.Equal(b, zeroBytes)
}

func Equal(a common.Address, b common.Address) bool {
	return a == b
}
