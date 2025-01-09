package cache

import (
	"sync"
	"time"
)

type V struct {
	Exp int64
	Fnc func()
}

type Time struct {
	dic map[string]V
	mut sync.Mutex
	now func() time.Time
	ttl time.Duration
}

func NewTime(ttl time.Duration) *Time {
	return &Time{
		dic: map[string]V{},
		mut: sync.Mutex{},
		now: func() time.Time { return time.Now() },
		ttl: ttl,
	}
}

func (t *Time) Delete(key string) {
	t.mut.Lock()
	delete(t.dic, key)
	t.mut.Unlock()
}

func (t *Time) Ensure(key string, fnc func()) {
	t.mut.Lock()
	t.dic[key] = V{Exp: t.now().Add(t.ttl).Unix(), Fnc: fnc}
	t.mut.Unlock()
}

func (t *Time) Expire(cyc time.Duration) {
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
func (t *Time) Length() int {
	t.mut.Lock()
	siz := len(t.dic)
	t.mut.Unlock()

	return siz
}
