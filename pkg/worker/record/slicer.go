package record

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/worker/object"
	"github.com/ethereum/go-ethereum/common"
)

const (
	// Max is the maximum wait duration allowed for packets to be delayed.
	Max = 1 * time.Minute
	// Wai is the multiplier added to the wait duration upon packet retry.
	Wai = 5 * time.Second
)

type SlicerConfig struct {
	Max time.Duration
	Wai time.Duration
}

type Slicer struct {
	lis []*Record
	max time.Duration
	wai time.Duration
}

func NewSlicer(c SlicerConfig) *Slicer {
	if c.Max == 0 {
		c.Max = Max
	}
	if c.Wai == 0 {
		c.Wai = Wai
	}

	return &Slicer{
		lis: []*Record{},
		max: c.Max,
		wai: c.Wai,
	}
}

func (s *Slicer) Add() {
	// Prepare some of the required information from the latest record, or define
	// some defaults if we are adding the first record to this slicer. The
	// following defaults apply.
	//
	//     the default status is "unknown"
	//     the default time is the current time
	//     the default try is 1
	//     the default transaction is the zero hash
	//     the default wait starts with increments of 5 seconds, capped at 1 minute
	//

	var sta Status
	var tim time.Time
	var try int
	var txn common.Hash
	var wai time.Duration

	if len(s.lis) == 0 {
		try = 1
	} else {
		try = s.Try().Get() + 1
		txn = s.Txn().Get()
	}

	{
		sta = Unknown
		tim = time.Now().UTC()
		wai = minDur(time.Duration(try)*s.wai, s.max)
	}

	// Once the new record is added to this slicer, we modify it by reference in
	// the block below.

	{
		s.lis = append(s.lis, New())
	}

	{
		s.Sta().Set(sta)
		s.Tim().Set(tim)
		s.Try().Set(try)
		s.Txn().Set(txn)
		s.Wai().Set(wai)
	}
}

func (s *Slicer) Get(ind int) *Record {
	return s.lis[ind]
}

func (s *Slicer) Len() int {
	return len(s.lis)
}

func (s *Slicer) Prv() *Record {
	return s.lis[len(s.lis)-2]
}

func (s *Slicer) Err() object.Interface[error] {
	return s.lis[len(s.lis)-1].Err()
}

func (s *Slicer) Sta() object.Interface[Status] {
	return s.lis[len(s.lis)-1].Sta()
}

func (s *Slicer) Tim() object.Interface[time.Time] {
	return s.lis[len(s.lis)-1].Tim()
}

func (s *Slicer) Try() object.Interface[int] {
	return s.lis[len(s.lis)-1].Try()
}

func (s *Slicer) Txn() object.Interface[common.Hash] {
	return s.lis[len(s.lis)-1].Txn()
}

func (s *Slicer) Wai() object.Interface[time.Duration] {
	return s.lis[len(s.lis)-1].Wai()
}

func minDur(a time.Duration, b time.Duration) time.Duration {
	if a < b {
		return a
	}

	return b
}
