package engine

import (
	"github.com/anubis-game/apiserver/pkg/router"
)

type Turn struct {
	Qdr byte
	Agl byte
}

// turn stores the desired travel direction for the given byte ID. The method
// strictly requires to be called sequentially in order to function properly.
// Further a note on an implementation detail. The *Engine type starts out with
// a preallocated empty slice of Turn structs. The game starts with every player
// having a current direction, but no desired direction. As long as the player
// does not provide their own desired change in direction, the player keeps
// moving in the same direction.
func (e *Engine) turn(pac router.Packet) {
	e.tur[pac.Uid].Qdr = pac.Byt[0]
	e.tur[pac.Uid].Agl = pac.Byt[1]
}
