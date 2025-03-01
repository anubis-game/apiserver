package filler

import (
	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/vector"
)

type Interface interface {
	Daemon()
	Energy(siz byte) *energy.Energy
	Vector(uid byte) *vector.Vector
}
