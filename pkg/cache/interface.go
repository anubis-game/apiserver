package cache

type Interface[T any] interface {
	//
	Create(key string, val T) bool

	//
	Delete(key string)

	//
	Escape(key string) T

	//
	Exists(key string) bool

	//
	Length() int

	//
	Ranger(fnc func(key string, val T))

	//
	Search(key string) T

	//
	Update(key string, val T)
}

type Testing interface {
	//
	Fatal(args ...any)
}
