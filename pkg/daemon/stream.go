package daemon

import (
	"github.com/anubis-game/apiserver/pkg/stream"
)

func (d *Daemon) Stream() *stream.Stream {
	return d.str
}
