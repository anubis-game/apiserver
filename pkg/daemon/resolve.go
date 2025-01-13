package daemon

import "github.com/anubis-game/apiserver/pkg/worker"

func (d *Daemon) Resolve() worker.Daemon {
	return d.res
}
