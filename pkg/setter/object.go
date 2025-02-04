package setter

type Setter[T comparable] struct {
	val T
}

func New[T comparable]() *Setter[T] {
	return &Setter[T]{}
}

func (o *Setter[T]) Emp() bool {
	var zer T
	return o.val == zer
}

func (o *Setter[T]) Get() T {
	return o.val
}

func (o *Setter[T]) Set(val T) {
	o.val = val
}
