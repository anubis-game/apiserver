package daemon

import "github.com/anubis-game/apiserver/pkg/engine"

func (d *Daemon) Engine() *engine.Engine {
	return d.eng
}
