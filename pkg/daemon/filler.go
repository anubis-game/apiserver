package daemon

import "github.com/anubis-game/apiserver/pkg/filler"

func (d *Daemon) Filler() *filler.Filler {
	return d.fil
}
