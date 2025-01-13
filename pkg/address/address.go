package address

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	zeroAddress common.Address
)

func Empty(a common.Address) bool {
	return a == zeroAddress
}

func Equal(a common.Address, b common.Address) bool {
	return a == b
}
