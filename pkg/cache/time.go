package cache

import (
	"sync"
	"time"
)

type V struct {
	Exp int64
	Fnc func()
}

type Time[K comparable] struct {
	dic map[K]V
	mut sync.RWMutex
	now func() time.Time
}

func NewTime[K comparable]() *Time[K] {
	return &Time[K]{
		dic: map[K]V{},
		mut: sync.RWMutex{},
		now: func() time.Time { return time.Now() },
	}
}

func (t *Time[K]) Delete(key K) {
	t.mut.Lock()
	delete(t.dic, key)
	t.mut.Unlock()
}

func (t *Time[K]) Ensure(key K, ttl time.Duration, fnc func()) {
	t.mut.Lock()
	t.dic[key] = V{Exp: t.now().Add(ttl).Unix(), Fnc: fnc}
	t.mut.Unlock()
}

func (t *Time[K]) Exists(key K) bool {
	t.mut.RLock()
	_, exi := t.dic[key]
	t.mut.RUnlock()

	return exi
}

func (t *Time[K]) Expire(cyc time.Duration) {
	var tic *time.Ticker
	{
		tic = time.NewTicker(cyc)
	}

	for {
		{
			<-tic.C
		}

		var now int64
		{
			now = t.now().Unix()
		}

		t.mut.Lock()
		for k, v := range t.dic {
			if now > v.Exp {
				v.Fnc()
				delete(t.dic, k)
			}
		}
		t.mut.Unlock()
	}
}
