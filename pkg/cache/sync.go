package cache

import "sync"

// Sync is a typed cache implementation, leveraging sync.Map for concurrency safety.
type Sync[T any] struct {
	dic sync.Map
}

func NewSync[T any]() Interface[T] {
	return &Sync[T]{}
}

// Create stores the given key-value pair if it does not already exist and
// returns whether the provided key already existed. Create uses sync.Map's
// LoadOrStore.
func (d *Sync[T]) Create(key string, val T) bool {
	_, exi := d.dic.LoadOrStore(key, val)
	return exi
}

// Delete simply removes the given key from the typed cache. Delete uses a
// write-lock. Delete uses sync.Map's Delete.
func (d *Sync[T]) Delete(key string) {
	d.dic.Delete(key)
}

// Escape is a Search-and-Delete, returning the value of the deleted key. Escape
// uses a write-lock. Escape uses sync.Map's LoadAndDelete.
func (d *Sync[T]) Escape(key string) T {
	if val, ok := d.dic.LoadAndDelete(key); ok {
		return val.(T)
	}
	var zer T
	return zer
}

// Exists returns whether the given key is already set. Exists uses sync.Map's
// Load.
func (d *Sync[T]) Exists(key string) bool {
	_, exi := d.dic.Load(key)
	return exi
}

// Length returns the amount of key-value pairs currently maintained in the
// underlying cache. Length uses sync.Map's Range.
func (d *Sync[T]) Length() int {
	siz := 0

	d.dic.Range(func(_, _ any) bool {
		siz++
		return true
	})

	return siz
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache. Ranger uses sync.Map's Range.
func (d *Sync[T]) Ranger(fnc func(key string, val T)) {
	d.dic.Range(func(k, v any) bool {
		fnc(k.(string), v.(T))
		return true
	})
}

// Search returns the value of the given key, whether that key exists or not.
// Search uses a read-lock. Search uses sync.Map's Load.
func (d *Sync[T]) Search(key string) T {
	if v, ok := d.dic.Load(key); ok {
		return v.(T)
	}
	var zer T
	return zer
}

// Update sets the given key-value pair or overwrites it in case the given key
// existed before. Update uses sync.Map's Store.
func (d *Sync[T]) Update(key string, val T) {
	d.dic.Store(key, val)
}
