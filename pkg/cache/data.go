package cache

import "sync"

// Data is a typed cache implementation, optimized for simple read-write access
// patterns.
type Data[T any] struct {
	dic map[string]T
	mut sync.RWMutex
}

func NewData[T any]() Interface[T] {
	return &Data[T]{
		dic: map[string]T{},
		mut: sync.RWMutex{},
	}
}

// Create stores the given key-value pair if it does not already exist and
// returns whether the provided key already existed. Create uses a write-lock.
func (d *Data[T]) Create(key string, val T) bool {
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
func (d *Data[T]) Delete(key string) {
	d.mut.Lock()
	delete(d.dic, key)
	d.mut.Unlock()
}

// Escape is a Search-and-Delete, returning the value of the deleted key. Escape
// uses a write-lock.
func (d *Data[T]) Escape(key string) T {
	var old T

	d.mut.Lock()
	old = d.dic[key]
	delete(d.dic, key)
	d.mut.Unlock()

	return old
}

// Exists returns whether the given key is already set. Exists uses a read-lock.
func (d *Data[T]) Exists(key string) bool {
	d.mut.RLock()
	_, exi := d.dic[key]
	d.mut.RUnlock()

	return exi
}

// Length returns the amount of key-value pairs currently maintained in the
// underlying cache.
func (d *Data[T]) Length() int {
	d.mut.RLock()
	siz := len(d.dic)
	d.mut.RUnlock()

	return siz
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache. Ranger uses a read-lock.
func (d *Data[T]) Ranger(fnc func(key string, val T)) {
	d.mut.RLock()

	for k, v := range d.dic {
		fnc(k, v)
	}

	d.mut.RUnlock()
}

// Search returns the value of the given key, whether that key exists or not.
// Search uses a read-lock.
func (d *Data[T]) Search(key string) T {
	d.mut.RLock()
	val := d.dic[key]
	d.mut.RUnlock()

	return val
}

// Update sets the given key-value pair or overwrites it in case the given key
// existed before. Update uses a write-lock.
func (d *Data[T]) Update(key string, val T) {
	d.mut.Lock()
	d.dic[key] = val
	d.mut.Unlock()
}
