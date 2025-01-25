package daemon

import (
	"github.com/anubis-game/apiserver/pkg/server/handler/connect"
)

func (d *Daemon) Connect() *connect.Handler {
	return d.con
}
