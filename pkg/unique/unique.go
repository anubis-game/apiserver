package unique

import "encoding/binary"

type Unique[T comparable] struct {
	lis []T
	zer T
}

func New[T comparable](len int) *Unique[T] {
	return &Unique[T]{
		lis: make([]T, len),
		// zer is allocated once per instance
	}
}

func (u *Unique[T]) Create(val T) [2]byte {
	for i, x := range u.lis {
		if x == u.zer {
			u.lis[i] = val
			var id [2]byte
			binary.BigEndian.PutUint16(id[:], uint16(i))
			return id
		}
	}

	return [2]byte{}
}

func (u *Unique[T]) Delete(val T) {
	for i, x := range u.lis {
		if x == val {
			u.lis[i] = u.zer
			return
		}
	}
}
