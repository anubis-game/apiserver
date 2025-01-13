package cache

import (
	"runtime"

	"github.com/puzpuzpuz/xsync/v3"
)

// Pool is a typed cache implementation, leveraging xsync.MapOf for concurrency
// safety, running multiple goroutines concurrently to execute the given Ranger
// callback.
type Pool[K comparable, V any] struct {
	dic *xsync.MapOf[K, V]
	sem chan struct{}
}

func NewPool[K comparable, V any]() Interface[K, V] {
	return &Pool[K, V]{
		dic: xsync.NewMapOf[K, V](),
		sem: make(chan struct{}, runtime.NumCPU()*4),
	}
}

// Create stores the given key-value pair if it does not already exist and
// returns whether the provided key already existed. Create uses xsync.MapOf's
// LoadOrStore.
func (p *Pool[K, V]) Create(key K, val V) bool {
	_, exi := p.dic.LoadOrStore(key, val)
	return exi
}

// Delete simply removes the given key from the typed cache. Delete uses a
// write-lock. Delete uses xsync.MapOf's Delete.
func (p *Pool[K, V]) Delete(key K) {
	p.dic.Delete(key)
}

// Escape is a Search-and-Delete, returning the value of the deleted key. Escape
// uses a write-lock. Escape uses xsync.MapOf's LoadAndDelete.
func (p *Pool[K, V]) Escape(key K) V {
	val, _ := p.dic.LoadAndDelete(key)
	return val
}

// Exists returns whether the given key is already set. Exists uses
// xsync.MapOf's Load.
func (p *Pool[K, V]) Exists(key K) bool {
	_, exi := p.dic.Load(key)
	return exi
}

// Length returns the amount of key-value pairs currently maintained in the
// underlying cache. Length uses xsync.MapOf's Size.
func (p *Pool[K, V]) Length() int {
	return p.dic.Size()
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache. Ranger uses xsync.MapOf's Range. The given callback is executed
// concurrently using a semaphore pattern.
func (p *Pool[K, V]) Ranger(fnc func(K, V)) {
	p.dic.Range(func(k K, v V) bool {
		// The semaphore controls the amount of workers that are allowed to process
		// the given callback concurrently. For every iteration, we push a ticket
		// into the semaphore before doing the actual work.
		{
			p.sem <- struct{}{}
		}

		// A new goroutine is created for every piece of work. That way we can
		// execute the given callback concurrently. Note that the given key-value
		// pairs must be injected into the worker goroutine as arguments, in order
		// to work on the exact key-value pair that this iteration received in this
		// asynchronous environment.
		go func(key K, val V) {
			// Forward the current key-value pair to the provided callback and wait
			// for the work to be done.
			{
				fnc(key, val)
			}

			// Ensure we remove our ticket from the semaphore once all work was
			// completed. This frees up the occupied worker goroutine so that another
			// piece of work can be done concurrently.
			{
				<-p.sem
			}
		}(k, v)

		return true
	})
}

// Search returns the value of the given key, whether that key exists or not.
// Search uses a read-lock. Search uses xsync.MapOf's Load.
func (p *Pool[K, V]) Search(key K) (V, bool) {
	return p.dic.Load(key)
}

// Update sets the given key-value pair or overwrites it in case the given key
// existed before. Update uses xsync.MapOf's Store.
func (p *Pool[K, V]) Update(key K, val V) {
	p.dic.Store(key, val)
}
