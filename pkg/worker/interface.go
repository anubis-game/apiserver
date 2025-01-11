package worker

type Create[T any] interface {
	Create(T)
}

type Daemon interface {
	Daemon()
}

type Ensure[T any] interface {
	Ensure(T) (T, bool)
}
