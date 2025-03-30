package engine

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/random"
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
			Crd: &random.Random{},
			Cap: cap,
			Don: make(<-chan struct{}),
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

func tesNrg(e *Engine, c matrix.Coordinate) {
	var n []byte
	{
		n = energy.New(energy.Config{
			Crd: c,
			Siz: 5,
			Typ: 0,
		})
	}

	{
		p := c.Pt1()

		m, x := e.lkp.nrg[p]
		if !x {
			m = map[matrix.Coordinate]struct{}{c: {}}
		} else {
			m[c] = struct{}{}
		}

		e.lkp.nrg[p] = m
	}

	{
		e.mem.nrg[c] = n
	}
}

func tesVec(u byte) *vector.Vector {
	var qdr byte
	var agl byte
	{
		qdr = byte(1)
		agl = byte(117)
	}

	// Create a new Vector instance using a random head node.

	var vec *vector.Vector
	{
		vec = vector.New(vector.Config{
			Hea: matrix.Coordinate{
				X: 3000,
				Y: 4000,
			},
			Mot: vector.Motion{
				Qdr: qdr,
				Agl: agl,
			},
			Uid: u,
		})
	}

	// We initialize the head of the new Vector above with a single coordinate
	// object. Below we use this head node as basis for the Vector's expansion. 1
	// head plus 4 expansions gives us a Vector with 5 nodes, all lined up towards
	// the same direction, because we use the same motion configuration every
	// time.

	for range 4 {
		vec.Update(int(vector.Si/vector.Li), qdr, agl, vector.Nrm)
	}

	return vec
}

func tesWal(u byte) common.Address {
	return common.HexToAddress(fmt.Sprintf("0x%040d", u))
}
