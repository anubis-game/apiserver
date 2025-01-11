package cache

import (
	"github.com/puzpuzpuz/xsync/v3"
)

// Sxnc is a typed cache implementation, leveraging xsync.MapOf for concurrency
// safety.
type Sxnc[T any] struct {
	dic *xsync.MapOf[string, T]
}

func NewSxnc[T any]() Interface[T] {
	return &Sxnc[T]{
		dic: xsync.NewMapOf[string, T](),
	}
}

// Create stores the given key-value pair if it does not already exist and
// returns whether the provided key already existed. Create uses xsync.MapOf's
// LoadOrStore.
func (d *Sxnc[T]) Create(key string, val T) bool {
	_, exi := d.dic.LoadOrStore(key, val)
	return exi
}

// Delete simply removes the given key from the typed cache. Delete uses a
// write-lock. Delete uses xsync.MapOf's Delete.
func (d *Sxnc[T]) Delete(key string) {
	d.dic.Delete(key)
}

// Escape is a Search-and-Delete, returning the value of the deleted key. Escape
// uses a write-lock. Escape uses xsync.MapOf's LoadAndDelete.
func (d *Sxnc[T]) Escape(key string) T {
	val, _ := d.dic.LoadAndDelete(key)
	return val
}

// Exists returns whether the given key is already set. Exists uses
// xsync.MapOf's Load.
func (d *Sxnc[T]) Exists(key string) bool {
	_, exi := d.dic.Load(key)
	return exi
}

// Length returns the amount of key-value pairs currently maintained in the
// underlying cache. Length uses xsync.MapOf's Size.
func (d *Sxnc[T]) Length() int {
	return d.dic.Size()
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache. Ranger uses xsync.MapOf's Range.
func (d *Sxnc[T]) Ranger(fnc func(key string, val T)) {
	d.dic.Range(func(k string, v T) bool {
		fnc(k, v)
		return true
	})
}

// Search returns the value of the given key, whether that key exists or not.
// Search uses a read-lock. Search uses xsync.MapOf's Load.
func (d *Sxnc[T]) Search(key string) T {
	val, _ := d.dic.Load(key)
	return val
}

// Update sets the given key-value pair or overwrites it in case the given key
// existed before. Update uses xsync.MapOf's Store.
func (d *Sxnc[T]) Update(key string, val T) {
	d.dic.Store(key, val)
}
