package worker

type Daemon interface {
	Daemon()
}

type Create[T any] interface {
	Create(T)
}

type Ensure[T any] interface {
	Ensure(T) (T, bool)
}
