package engine

import "github.com/anubis-game/apiserver/pkg/router"

func (e *Engine) delete(pac router.Packet) {
	delete(e.cli, pac.Cli.Wallet())
}
