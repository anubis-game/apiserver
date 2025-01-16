package cache

import (
	"github.com/puzpuzpuz/xsync/v3"
)

// Sxnc is a typed cache implementation, leveraging xsync.MapOf for concurrency
// safety.
type Sxnc[K comparable, V any] struct {
	dic *xsync.MapOf[K, V]
}

func NewSxnc[K comparable, V any]() Interface[K, V] {
	return &Sxnc[K, V]{
		dic: xsync.NewMapOf[K, V](),
	}
}

// Create stores the given key-value pair if it does not already exist and
// returns whether the provided key already existed. Create uses xsync.MapOf's
// LoadOrStore.
func (d *Sxnc[K, V]) Create(key K, val V) bool {
	_, exi := d.dic.LoadOrStore(key, val)
	return exi
}

// Delete simply removes the given key from the typed cache. Delete uses a
// write-lock. Delete uses xsync.MapOf's Delete.
func (d *Sxnc[K, V]) Delete(key K) {
	d.dic.Delete(key)
}

// Escape is a Search-and-Delete, returning the value of the deleted key. Escape
// uses a write-lock. Escape uses xsync.MapOf's LoadAndDelete.
func (d *Sxnc[K, V]) Escape(key K) V {
	val, _ := d.dic.LoadAndDelete(key)
	return val
}

// Exists returns whether the given key is already set. Exists uses
// xsync.MapOf's Load.
func (d *Sxnc[K, V]) Exists(key K) bool {
	_, exi := d.dic.Load(key)
	return exi
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache. Ranger uses xsync.MapOf's Range.
func (d *Sxnc[K, V]) Ranger(fnc func(K, V)) {
	d.dic.Range(func(k K, v V) bool {
		fnc(k, v)
		return true
	})
}

// Search returns the value of the given key, whether that key exists or not.
// Search uses a read-lock. Search uses xsync.MapOf's Load.
func (d *Sxnc[K, V]) Search(key K) (V, bool) {
	return d.dic.Load(key)
}

// Update sets the given key-value pair or overwrites it in case the given key
// existed before. Update uses xsync.MapOf's Store.
func (d *Sxnc[K, V]) Update(key K, val V) {
	d.dic.Store(key, val)
}

// length returns the amount of key-value pairs currently maintained in the
// underlying cache. length uses xsync.MapOf's Size.
func (d *Sxnc[K, V]) length() int { // nolint:unused
	return d.dic.Size()
}
