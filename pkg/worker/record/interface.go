package record

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/worker/object"
	"github.com/ethereum/go-ethereum/common"
)

type Interface interface {
	Add()
	Get(int) *Record
	Len() int
	Prv() *Record

	Err() object.Interface[error]
	Sta() object.Interface[Status]
	Tim() object.Interface[time.Time]
	Try() object.Interface[int]
	Txn() object.Interface[common.Hash]
	Wai() object.Interface[time.Duration]
}
