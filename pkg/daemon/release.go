package daemon

import "github.com/anubis-game/apiserver/pkg/worker"

func (d *Daemon) Release() worker.Daemon {
	return d.rel
}
