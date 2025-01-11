package worker

type Daemon interface {
	Daemon()
}

type Ensure[T any] interface {
	Ensure(T)
}

type Router[T any] interface {
	Router(T) (T, bool)
}
