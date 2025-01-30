package daemon

import "github.com/anubis-game/apiserver/pkg/worker"

func (d *Daemon) Worker() worker.Daemon {
	return d.wrk
}
