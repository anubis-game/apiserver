package engine

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
)

func tesEng(cap int) *Engine {
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
			Fil: &testFiller{},
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

type testFiller struct{}

func (f *testFiller) Daemon() {}

func (f *testFiller) Energy(siz byte) *energy.Energy { return nil }

func (f *testFiller) Vector() *vector.Vector {
	var mot vector.Motion
	{
		mot = vector.Motion{
			Qdr: byte(2),
			Agl: byte(128),
		}
	}

	var vec *vector.Vector
	{
		vec = vector.New(vector.Config{
			Hea: matrix.Coordinate{
				X: 1000,
				Y: 1000,
			},
			Mot: mot,
		})
	}

	for range 4 {
		vec.Update(int(vector.Si/vector.Li), mot.Qdr, mot.Agl, vector.Nrm)
	}

	return vec
}
