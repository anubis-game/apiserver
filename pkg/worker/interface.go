package worker

import "time"

type Create[P any] interface {
	Create(P)
}

type Daemon interface {
	Daemon()
}

type Ensure[K comparable, P any] interface {
	Ensure(P) (P, K, time.Duration)
}
