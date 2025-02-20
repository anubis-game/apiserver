package unique

import (
	"sync"
)

type Unique[K comparable, V Number] struct {
	lis []K
	mut sync.Mutex
	rev map[K]V
	zer K
}

func New[K comparable, V Number]() *Unique[K, V] {
	var len int
	{
		len = maxOf[V]()
	}

	return &Unique[K, V]{
		lis: make([]K, len),
		rev: make(map[K]V, len),
	}
}

func (u *Unique[K, V]) Delete(v K) {
	u.mut.Lock()

	i, e := u.rev[v]
	if e {
		delete(u.rev, v)
		u.lis[i] = u.zer
	}

	u.mut.Unlock()
}

func (u *Unique[K, V]) Ensure(v K) V {
	u.mut.Lock()

	i, e := u.rev[v]
	if e {
		u.mut.Unlock()

		return i
	}

	for i, x := range u.lis {
		if x == u.zer {
			u.lis[i] = v
			u.rev[v] = V(i)

			u.mut.Unlock()

			return V(i)
		}
	}

	u.mut.Unlock()
	return 0
}

// length is only used for testing.
func (u *Unique[K, V]) length() int {
	u.mut.Lock()
	l := len(u.rev)
	u.mut.Unlock()
	return l
}
