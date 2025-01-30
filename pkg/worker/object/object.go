package object

type Object[T comparable] struct {
	val T
}

func New[T comparable]() *Object[T] {
	return &Object[T]{}
}

func (o *Object[T]) Emp() bool {
	var zer T
	return o.val == zer
}

func (o *Object[T]) Get() T {
	return o.val
}

func (o *Object[T]) Set(val T) {
	o.val = val
}
