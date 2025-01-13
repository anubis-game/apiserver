package resolve

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/objectid"
)

type Packet struct {
	Kill        objectid.ID
	Winner      common.Address
	Loser       common.Address
	Timeout     time.Duration
	Transaction common.Hash
}
