package engine

import (
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/window"
	"github.com/ethereum/go-ethereum/common"
)

func (e *Engine) create(pac router.Packet) {
	// Upon joining, we add the user to the broadcasting worker pool in order to
	// provide them with realtime data primitives over the given client
	// connection. We check whether the given Wallet address is already part of
	// the broadcasting worker pool. If it is, then we do not want to do
	// unnecessary work, and instead return early.

	var wal common.Address
	{
		wal = pac.Cli.Wallet()
	}

	{
		_, exi := e.cli[wal]
		if exi {
			e.log.Log(
				"level", "warning",
				"message", "already joined",
				"wallet", wal.String(),
			)

			return
		}
	}

	{
		e.cli[wal] = pac.Cli
	}

	var win *window.Window
	{
		win = pac.Cli.Window()
	}

	// Put the player randomly onto the game map.
	{
		pac.Cli.Stream(schema.Encode(schema.Join, win.Bytes()))
	}

	// TODO we need to stream all relevant map details for the initial view
}
