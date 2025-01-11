package cache

type Interface[K comparable, V any] interface {
	//
	Create(K, V) bool

	//
	Delete(K)

	//
	Escape(K) V

	//
	Exists(K) bool

	//
	Length() int

	//
	Ranger(fnc func(K, V))

	//
	Search(K) (V, bool)

	//
	Update(K, V)
}

type Testing interface {
	//
	Fatal(args ...any)
}
