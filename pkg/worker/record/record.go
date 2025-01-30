package record

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/worker/object"
	"github.com/ethereum/go-ethereum/common"
)

type Record struct {
	err object.Interface[error]
	sta object.Interface[Status]
	tim object.Interface[time.Time]
	try object.Interface[int]
	txn object.Interface[common.Hash]
	wai object.Interface[time.Duration]
}

func New() *Record {
	return &Record{
		err: object.New[error](),
		sta: object.New[Status](),
		tim: object.New[time.Time](),
		try: object.New[int](),
		txn: object.New[common.Hash](),
		wai: object.New[time.Duration](),
	}
}

func (r *Record) Err() object.Interface[error] {
	return r.err
}

func (r *Record) Sta() object.Interface[Status] {
	return r.sta
}

func (r *Record) Tim() object.Interface[time.Time] {
	return r.tim
}

func (r *Record) Try() object.Interface[int] {
	return r.try
}

func (r *Record) Txn() object.Interface[common.Hash] {
	return r.txn
}

func (r *Record) Wai() object.Interface[time.Duration] {
	return r.wai
}
