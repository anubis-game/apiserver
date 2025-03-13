package engine

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
)

func tesEng(cap int) *Engine {
	var fil *filler.Filler
	{
		fil = filler.New(filler.Config{
			Cap: cap,
			Don: make(<-chan struct{}),
			Log: logger.Fake(),
		})
	}

	{
		go fil.Daemon()
	}

	var rtr *router.Router
	{
		rtr = router.New(router.Config{
			Cap: cap,
		})
	}

	var eng *Engine
	{
		eng = New(Config{
			Cap: cap,
			Don: make(<-chan struct{}),
			Fil: fil,
			Log: logger.Fake(),
			Rtr: rtr.Engine(),
			Tkx: tokenx.New[common.Address](),
			Uni: unique.New[common.Address, byte](),
			Wrk: &worker.Worker{},
		})
	}

	for u := range byte(cap) {
		eng.uni.Ensure(tesWal(u))
	}

	return eng
}

func tesWal(u byte) common.Address {
	return common.HexToAddress(fmt.Sprintf("0x%040d", u))
}
