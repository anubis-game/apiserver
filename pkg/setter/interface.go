package setter

type Interface[T comparable] interface {
	Emp() bool
	Get() T
	Set(val T)
}
