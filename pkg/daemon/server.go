package daemon

import (
	"github.com/anubis-game/apiserver/pkg/server"
)

func (d *Daemon) Server() *server.Server {
	return server.New(server.Config{
		Env: d.env,
		Lis: d.lis,
		Log: d.log,
		Rtr: d.rtr,
		Str: d.str,
	})
}
