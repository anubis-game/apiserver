package unique

import (
	"sync"
)

// Unique generates numerical IDs in constant time within a given capacity. Note
// that we use a normal sync.Mutex for synchronization instead of a
// sync.RWMutex, because the synchronized code is many times faster than the
// additional overhead incurred by sync.RWMutex.
type Unique[K comparable, V Number] struct {
	ind int
	lis []V
	mut sync.Mutex
	rev map[K]V
}

func New[K comparable, V Number]() *Unique[K, V] {
	var len int
	{
		len = length[V]()
	}

	// The mechanism of this stack based ID generation does only work if there is
	// a stack that we can draw from. So we have to initialize the stack values.

	var lis []V
	{
		lis = make([]V, len)
	}

	for i := range lis {
		lis[i] = V(i)
	}

	return &Unique[K, V]{
		lis: lis,
		rev: make(map[K]V, len),
	}
}

func (u *Unique[K, V]) Delete(k K) {
	u.mut.Lock()

	r, e := u.rev[k]
	if e {
		delete(u.rev, k)
		u.ind--
		u.lis[u.ind] = r
	}

	u.mut.Unlock()
}

func (u *Unique[K, V]) Ensure(k K) V {
	u.mut.Lock()

	// Returning any allocated value early guarantees idempotency.

	r, e := u.rev[k]
	if e {
		u.mut.Unlock()
		return r
	}

	// If we run out of capacity we stop early

	if u.ind >= len(u.lis) {
		u.mut.Unlock()
		return 0
	}

	// We just take the next available item from the stack and increment our
	// pointer.

	v := u.lis[u.ind]
	u.rev[k] = v
	u.ind++
	u.mut.Unlock()

	return v
}

func (u *Unique[K, V]) Exists(k K) bool {
	u.mut.Lock()
	_, e := u.rev[k]
	u.mut.Unlock()
	return e
}

// Length provides the amount of currently allocated IDs. Since ID allocation is
// strictly sequential, all IDs can be iterated in an idomatic range loop. The
// iterated order does not reflect the order of creation nor deletion.
//
//	for x := range u.Length() {
//	  ...
//	}
func (u *Unique[K, V]) Length() V {
	u.mut.Lock()
	l := u.ind
	u.mut.Unlock()
	return V(l)
}
