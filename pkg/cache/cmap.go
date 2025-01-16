package cache

// Cmap is a typed cache implementation using channels for synchronization.
type Cmap[K comparable, V any] struct {
	chn chan struct{}
	dic map[K]V
}

func NewCmap[K comparable, V any]() Interface[K, V] {
	return &Cmap[K, V]{
		chn: make(chan struct{}, 1),
		dic: make(map[K]V),
	}
}

// Create stores the given key-value pair if it does not already exist and
// returns whether the provided key already existed. Create uses a write-lock.
func (d *Cmap[K, V]) Create(key K, val V) bool {
	var exi bool

	d.chn <- struct{}{}
	_, exi = d.dic[key]
	if !exi {
		d.dic[key] = val
	}
	<-d.chn

	return exi
}

// Delete simply removes the given key from the typed cache. Delete uses a
// write-lock.
func (d *Cmap[K, V]) Delete(key K) {
	d.chn <- struct{}{}
	delete(d.dic, key)
	<-d.chn
}

// Escape is a Search-and-Delete, returning the value of the deleted key. Escape
// uses a write-lock.
func (d *Cmap[K, V]) Escape(key K) V {
	var old V

	d.chn <- struct{}{}
	old = d.dic[key]
	delete(d.dic, key)
	<-d.chn

	return old
}

// Exists returns whether the given key is already set. Exists uses a read-lock.
func (d *Cmap[K, V]) Exists(key K) bool {
	d.chn <- struct{}{}
	_, exi := d.dic[key]
	<-d.chn

	return exi
}

// Ranger executes the given callback for every key-value pair in the underlying
// cache. Ranger uses a read-lock.
func (d *Cmap[K, V]) Ranger(fnc func(K, V)) {
	d.chn <- struct{}{}

	for k, v := range d.dic {
		fnc(k, v)
	}

	<-d.chn
}

// Search returns the value of the given key, whether that key exists or not.
// Search uses a read-lock.
func (d *Cmap[K, V]) Search(key K) (V, bool) {
	d.chn <- struct{}{}
	val, exi := d.dic[key]
	<-d.chn

	return val, exi
}

// Update sets the given key-value pair or overwrites it in case the given key
// existed before. Update uses a write-lock.
func (d *Cmap[K, V]) Update(key K, val V) {
	d.chn <- struct{}{}
	d.dic[key] = val
	<-d.chn
}

// length returns the amount of key-value pairs currently maintained in the
// underlying cache.
func (d *Cmap[K, V]) length() int { // nolint:unused
	d.chn <- struct{}{}
	siz := len(d.dic)
	<-d.chn

	return siz
}
