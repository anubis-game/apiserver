package unique

import (
	"encoding/binary"
	"sync"
)

type Unique[T comparable] struct {
	lis []T
	mut sync.Mutex
	rev map[T]uint16
	zer T
}

func New[T comparable](len int) *Unique[T] {
	return &Unique[T]{
		lis: make([]T, len),
		rev: make(map[T]uint16, len),
	}
}

func (u *Unique[T]) Delete(v T) {
	u.mut.Lock()

	i, e := u.rev[v]
	if e {
		delete(u.rev, v)
		u.lis[i] = u.zer
	}

	u.mut.Unlock()
}

func (u *Unique[T]) Ensure(v T) [2]byte {
	u.mut.Lock()

	i, e := u.rev[v]
	if e {
		u.mut.Unlock()

		var b [2]byte
		binary.BigEndian.PutUint16(b[:], uint16(i))
		return b
	}

	for i, x := range u.lis {
		if x == u.zer {
			u.lis[i] = v
			u.rev[v] = uint16(i)

			u.mut.Unlock()

			var b [2]byte
			binary.BigEndian.PutUint16(b[:], uint16(i))
			return b
		}
	}

	u.mut.Unlock()
	return [2]byte{}
}

// length is only used for testing.
func (u *Unique[T]) length() int {
	u.mut.Lock()
	l := len(u.rev)
	u.mut.Unlock()
	return l
}
