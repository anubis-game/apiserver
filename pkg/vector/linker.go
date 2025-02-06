package vector

import "github.com/anubis-game/apiserver/pkg/object"

type Linker struct {
	val object.Object
	nxt *Linker
}
