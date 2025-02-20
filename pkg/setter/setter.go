package setter

type Setter[T any] struct {
	val T
}

func New[T any]() *Setter[T] {
	return &Setter[T]{}
}

func (o *Setter[T]) Get() T {
	return o.val
}

func (o *Setter[T]) Set(val T) {
	o.val = val
}
