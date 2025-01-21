package daemon

import "github.com/anubis-game/apiserver/pkg/random"

func (d *Daemon) Random() *random.Random {
	return d.ran
}
