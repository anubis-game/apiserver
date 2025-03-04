package vector

import "github.com/anubis-game/apiserver/pkg/matrix"

// Linker is a single pointer reference within the Vector's linked list. We
// utilize a linked list for constant time updates of the Vector's head and
// tail.
type Linker struct {
	crd matrix.Coordinate
	// hid tracks the amount of hidden nodes that any given node is accounting for
	// between itself and its previous node. The current head node tracks hidden
	// nodes between itself and the previous head. The tail node does not track
	// any hidden nodes.
	hid int8
	nxt *Linker
	prv *Linker
}
