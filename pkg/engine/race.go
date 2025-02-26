package engine

import (
	"github.com/anubis-game/apiserver/pkg/vector"
)

// race switches between normal and racing speed for the given byte ID. The
// method strictly requires to be called sequentially in order to function
// properly. Further a note on an implementation detail. The *Engine type starts
// out with a preallocated empty slice of racing bytes. The game starts with
// every player operating at normal speed. The first call to Engine.race() must
// therefore switch to racing mode.
func (e *Engine) race(uid byte) {
	if e.rac[uid] == vector.Rcn {
		e.rac[uid] = vector.Nrm
	} else {
		e.rac[uid] = vector.Rcn
	}
}
