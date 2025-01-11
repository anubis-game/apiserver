package cache

import "sync"

// TODO if address types perform better, adapt /stream logic

// Data is a typed cache implementation, optimized for simple read-write access
// patterns.
type Data[K comparable, V any] struct {
	dic map[K]V
	mut sync.RWMutex
}

func NewData[K comparable, V any]() Interface[K, V] {
	return &Data[K, V]{
		dic: map[K]V{},
		mut: sync.RWMutex{},
	}
}

// Create stores the given key-value pair if it does not already exist and
// returns whether the provided key already existed. Create uses a write-lock.
func (d *Data[K, V]) Create(key K, val V) bool {
	var exi bool

	d.mut.Lock()
	_, exi = d.dic[key]
	if !exi {
		d.dic[key] = val
	}
	d.mut.Unlock()

	return exi
}

// Delete simply removes the given key from the typed cache. Delete uses a
// write-lock.
func (d *Data[K, V]) Delete(key K) {
	d.mut.Lock()
	delete(d.dic, key)
	d.mut.Unlock()
}

// Escape is a Search-and-Delete, returning the value of the deleted key. Escape
// uses a write-lock.
func (d *Data[K, V]) Escape(key K) V {
	var old V

	d.mut.Lock()
	old = d.dic[key]
	delete(d.dic, key)
	d.mut.Unlock()

	return old
}

// Exists returns whether the given key is already set. Exists uses a read-lock.
func (d *Data[K, V]) Exists(key K) bool {
	d.mut.RLock()
	_, exi := d.dic[key]
	d.mut.RUnlock()

	return exi
}

// Length returns the amount of key-value pairs currently maintained in the
// underlying cache.
func (d *Data[K, V]) Length() int {
	d.mut.RLock()
	siz := len(d.dic)
	d.mut.RUnlock()

	return siz
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache. Ranger uses a read-lock.
func (d *Data[K, V]) Ranger(fnc func(K, V)) {
	d.mut.RLock()

	for k, v := range d.dic {
		fnc(k, v)
	}

	d.mut.RUnlock()
}

// Search returns the value of the given key, whether that key exists or not.
// Search uses a read-lock.
func (d *Data[K, V]) Search(key K) (V, bool) {
	d.mut.RLock()
	val, exi := d.dic[key]
	d.mut.RUnlock()

	return val, exi
}

// Update sets the given key-value pair or overwrites it in case the given key
// existed before. Update uses a write-lock.
func (d *Data[K, V]) Update(key K, val V) {
	d.mut.Lock()
	d.dic[key] = val
	d.mut.Unlock()
}
