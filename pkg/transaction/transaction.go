package transaction

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	zeroHash common.Hash
)

func Empty(a common.Hash) bool {
	return a == zeroHash
}

func Equal(a common.Hash, b common.Hash) bool {
	return a == b
}
