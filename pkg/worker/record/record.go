package record

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/setter"
	"github.com/ethereum/go-ethereum/common"
)

type Record struct {
	err setter.Interface[error]
	sta setter.Interface[Status]
	tim setter.Interface[time.Time]
	try setter.Interface[int]
	txn setter.Interface[common.Hash]
	wai setter.Interface[time.Duration]
}

func New() *Record {
	return &Record{
		err: setter.New[error](),
		sta: setter.New[Status](),
		tim: setter.New[time.Time](),
		try: setter.New[int](),
		txn: setter.New[common.Hash](),
		wai: setter.New[time.Duration](),
	}
}

func (r *Record) Err() setter.Interface[error] {
	return r.err
}

func (r *Record) Sta() setter.Interface[Status] {
	return r.sta
}

func (r *Record) Tim() setter.Interface[time.Time] {
	return r.tim
}

func (r *Record) Try() setter.Interface[int] {
	return r.try
}

func (r *Record) Txn() setter.Interface[common.Hash] {
	return r.txn
}

func (r *Record) Wai() setter.Interface[time.Duration] {
	return r.wai
}
