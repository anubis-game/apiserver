package vector

import "github.com/anubis-game/apiserver/pkg/object"

// Linker is a single pointer reference within the Vector's linked list. We
// utilize a linked list for constant time updates of the Vector's head and
// tail.
type Linker struct {
	val object.Object
	nxt *Linker
}
