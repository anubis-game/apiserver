package engine

import (
	"github.com/anubis-game/apiserver/pkg/router"
)

// turn stores the desired travel direction for the given byte ID. The method
// strictly requires to be called sequentially in order to function properly.
// Further a note on an implementation detail. The *Engine type starts out with
// a preallocated empty slice of Turn structs. The game starts with every player
// having a current direction, but no desired direction. As long as the player
// does not provide their own desired change in direction, the player keeps
// moving in the same direction.
func (e *Engine) turn(pac router.Turn) {
	e.ply.tur[pac.Uid].Qdr = pac.Qdr
	e.ply.tur[pac.Uid].Agl = pac.Agl
}
