package record

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/ethereum/go-ethereum/common"
)

type Interface interface {
	Add()
	Get(int) *Record
	Len() int
	Prv() *Record

	Err() setter.Interface[error]
	Sta() setter.Interface[Status]
	Tim() setter.Interface[time.Time]
	Try() setter.Interface[int]
	Txn() setter.Interface[common.Hash]
	Wai() setter.Interface[time.Duration]
}
