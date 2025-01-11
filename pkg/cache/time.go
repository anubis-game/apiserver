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
	mut sync.Mutex
	now func() time.Time
	ttl time.Duration
}

func NewTime[K comparable](ttl time.Duration) *Time[K] {
	return &Time[K]{
		dic: map[K]V{},
		mut: sync.Mutex{},
		now: func() time.Time { return time.Now() },
		ttl: ttl,
	}
}

func (t *Time[K]) Delete(key K) {
	t.mut.Lock()
	delete(t.dic, key)
	t.mut.Unlock()
}

func (t *Time[K]) Ensure(key K, fnc func()) {
	t.mut.Lock()
	t.dic[key] = V{Exp: t.now().Add(t.ttl).Unix(), Fnc: fnc}
	t.mut.Unlock()
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

// Length returns the amount of key-value pairs currently maintained in the
// underlying cache.
func (t *Time[K]) Length() int {
	t.mut.Lock()
	siz := len(t.dic)
	t.mut.Unlock()

	return siz
}
