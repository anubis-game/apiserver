package worker

type Create[P any] interface {
	Create(P)
}

type Daemon interface {
	Daemon()
}

type Ensure[P any] interface {
	Ensure(P) (P, bool)
}
