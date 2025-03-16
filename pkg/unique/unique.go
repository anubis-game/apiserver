package unique

import (
	"slices"
	"sync"
	"sync/atomic"

	"github.com/anubis-game/apiserver/pkg/number"
)

// Unique generates numerical IDs in constant time within a given capacity. Note
// that we use a normal sync.Mutex for synchronization instead of a
// sync.RWMutex, because the synchronized code is many times faster than the
// additional overhead incurred by sync.RWMutex.
type Unique[K comparable, V number.Number] struct {
	ind *atomic.Int32
	lis []V
	mut sync.Mutex
	rev []K
	zer K
}

func New[K comparable, V number.Number]() *Unique[K, V] {
	var len int
	{
		len = number.Length[V]()
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
		ind: &atomic.Int32{},
		lis: lis,
		rev: make([]K, len),
	}
}

func (u *Unique[K, V]) Delete(k K) {
	u.mut.Lock()

	for i, x := range u.rev {
		if x == k {
			u.lis[u.ind.Add(-1)] = V(i)
			u.rev[i] = u.zer
			break
		}
	}

	u.mut.Unlock()
}

func (u *Unique[K, V]) Ensure(k K) V {
	u.mut.Lock()

	// Returning any allocated value early guarantees idempotency.

	for i, x := range u.rev {
		if x == k {
			u.mut.Unlock()
			return V(i)
		}
	}

	// If we run out of capacity we stop early.

	i := u.ind.Load()
	if int(i) >= len(u.lis) {
		u.mut.Unlock()
		return 0
	}

	// We just take the next available item from the stack and increment our
	// pointer.

	v := u.lis[i]
	u.ind.Add(+1)
	u.rev[v] = k
	u.mut.Unlock()

	return v
}

func (u *Unique[K, V]) Exists(k K) bool {
	u.mut.Lock()

	if slices.Contains(u.rev, k) {
		u.mut.Unlock()
		return true
	}
	u.mut.Unlock()

	return false
}

// Length provides the amount of currently allocated IDs. Since ID allocation is
// strictly sequential, all IDs can be iterated in an idomatic range loop. The
// iterated order does not reflect the order of creation nor deletion.
//
//	for x := range u.Length() {
//	  ...
//	}
func (u *Unique[K, V]) Length() V {
	return V(u.ind.Load())
}
