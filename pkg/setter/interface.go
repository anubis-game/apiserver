package setter

type Interface[T any] interface {
	Get() T
	Set(val T)
}
