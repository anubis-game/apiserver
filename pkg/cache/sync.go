package cache

import "sync"

// Sync is a typed cache implementation, leveraging sync.Map for concurrency
// safety.
type Sync[K comparable, V any] struct {
	dic sync.Map
}

func NewSync[K comparable, V any]() Interface[K, V] {
	return &Sync[K, V]{}
}

// Create stores the given key-value pair if it does not already exist and
// returns whether the provided key already existed. Create uses sync.Map's
// LoadOrStore.
func (d *Sync[K, V]) Create(key K, val V) bool {
	_, exi := d.dic.LoadOrStore(key, val)
	return exi
}

// Delete simply removes the given key from the typed cache. Delete uses a
// write-lock. Delete uses sync.Map's Delete.
func (d *Sync[K, V]) Delete(key K) {
	d.dic.Delete(key)
}

// Escape is a Search-and-Delete, returning the value of the deleted key. Escape
// uses a write-lock. Escape uses sync.Map's LoadAndDelete.
func (d *Sync[K, V]) Escape(key K) V {
	val, exi := d.dic.LoadAndDelete(key)
	if exi {
		return val.(V)
	}

	var zer V
	return zer
}

// Exists returns whether the given key is already set. Exists uses sync.Map's
// Load.
func (d *Sync[K, V]) Exists(key K) bool {
	_, exi := d.dic.Load(key)
	return exi
}

// Length returns the amount of key-value pairs currently maintained in the
// underlying cache. Length uses sync.Map's Range.
func (d *Sync[K, V]) Length() int {
	siz := 0

	d.dic.Range(func(_ any, _ any) bool {
		siz++
		return true
	})

	return siz
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache. Ranger uses sync.Map's Range.
func (d *Sync[K, V]) Ranger(fnc func(K, V)) {
	d.dic.Range(func(k any, v any) bool {
		fnc(k.(K), v.(V))
		return true
	})
}

// Search returns the value of the given key, whether that key exists or not.
// Search uses a read-lock. Search uses sync.Map's Load.
func (d *Sync[K, V]) Search(key K) (V, bool) {
	val, exi := d.dic.Load(key)
	if exi {
		return val.(V), exi
	}

	var zer V
	return zer, exi
}

// Update sets the given key-value pair or overwrites it in case the given key
// existed before. Update uses sync.Map's Store.
func (d *Sync[K, V]) Update(key K, val V) {
	d.dic.Store(key, val)
}
