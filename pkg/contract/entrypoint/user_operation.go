package entrypoint

import (
	"github.com/ethereum/go-ethereum/common"
)

type UserOperation struct {
	Sender   common.Address
	CallData []byte
}
