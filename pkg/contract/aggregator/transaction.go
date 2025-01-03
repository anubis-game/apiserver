package aggregator

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	Target   common.Address
	Value    *big.Int
	CallData []byte
}
