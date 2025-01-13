package release

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Packet struct {
	Loser       common.Address
	Timeout     time.Duration
	Transaction common.Hash
}
