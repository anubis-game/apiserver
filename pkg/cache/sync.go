package cache

import "sync"

// Sync is a typed cache implementation, leveraging sync.Map for concurrency safety.
type Sync[T any] struct {
	dic sync.Map
}

func NewSync[T any]() Interface[T] {
	return &Sync[T]{}
}

// Create sets the given key-value pair and returns whether the provided key
// already existed. Uses sync.Map's LoadOrStore.
func (d *Sync[T]) Create(key string, val T) bool {
	_, loaded := d.dic.LoadOrStore(key, val)
	return loaded
}

// Delete removes the given key from the typed cache.
func (d *Sync[T]) Delete(key string) {
	d.dic.Delete(key)
}

// Escape is a Search-and-Delete, returning the value of the deleted key.
func (d *Sync[T]) Escape(key string) T {
	if val, ok := d.dic.LoadAndDelete(key); ok {
		return val.(T)
	}
	var zero T
	return zero
}

// Exists returns whether the given key is already set.
func (d *Sync[T]) Exists(key string) bool {
	_, exists := d.dic.Load(key)
	return exists
}

// Length returns the number of key-value pairs in the underlying cache.
func (d *Sync[T]) Length() int {
	count := 0
	d.dic.Range(func(_, _ any) bool {
		count++
		return true
	})

	return count
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache.
func (d *Sync[T]) Ranger(fnc func(key string, val T)) {
	d.dic.Range(func(k, v any) bool {
		fnc(k.(string), v.(T))
		return true
	})
}

// Search returns the value of the given key, whether that key exists or not.
func (d *Sync[T]) Search(key string) T {
	if v, ok := d.dic.Load(key); ok {
		return v.(T)
	}
	var zero T
	return zero
}

// Update sets the given key-value pair, overwriting it if the key existed before.
func (d *Sync[T]) Update(key string, val T) {
	d.dic.Store(key, val)
}
