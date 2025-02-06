package test

import "github.com/anubis-game/apiserver/pkg/object"

type Mpbody struct {
	a uint16
	z uint16
	m map[uint16]object.Object
}

func (b *Mpbody) Move() {
	b.m[b.a+1] = b.m[b.z]

	delete(b.m, b.z)

	b.a++
	b.z++
}

type Scbody struct {
	m []object.Object
}

func (b *Scbody) Move() {
	las := len(b.m) - 1

	tai := b.m[las] // get the tail

	copy(b.m[1:], b.m[:las])

	b.m[0] = tai // tail becomes head
}
