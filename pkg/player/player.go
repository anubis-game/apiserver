package player

import (
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/vector"
)

type Player struct {
	Cli *client.Client
	Uid [2]byte
	Vec *vector.Vector
}
