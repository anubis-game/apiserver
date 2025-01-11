package daemon

import (
	"github.com/anubis-game/apiserver/pkg/server"
)

func (d *Daemon) Server() *server.Server {
	return d.ser
}
