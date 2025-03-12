package engine

import (
	"fmt"
	"time"

	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Cap int
	Don <-chan struct{}
	Fil filler.Interface
	Log logger.Interface
	Rtr *router.Engine
	Tkx *tokenx.TokenX[common.Address]
	Uni *unique.Unique[common.Address, byte]
	Wrk worker.Ensure
}

type Engine struct {
	// don is the global channel to signal program termination. If this channel is
	// closed, then all streaming connections should be terminated gracefully.
	don <-chan struct{}
	// filler
	fil filler.Interface
	// log is a simple logger interface to print system relevant information.
	log logger.Interface
	// rtr is the bridge synchronizing the server handler and the game engine
	rtr *router.Engine
	// tic is the global pointer keeping track of the fanout related time ticks.
	// This timestamp tracks at which point the latest fanout procedure has been
	// executed. The first tick is initialized in Engine.Daemon().
	tic time.Time
	// tkx
	tkx *tokenx.TokenX[common.Address]
	// uni
	uni *unique.Unique[common.Address, byte]
	// wrk
	wrk worker.Ensure

	// lkp
	lkp *lookup
	// mem
	mem *memory
	// ply
	ply *player
}

func New(c Config) *Engine {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Fil == nil {
		tracer.Panic(fmt.Errorf("%T.Fil must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}
	if c.Rtr == nil {
		tracer.Panic(fmt.Errorf("%T.Rtr must not be empty", c))
	}
	if c.Tkx == nil {
		tracer.Panic(fmt.Errorf("%T.Tkx must not be empty", c))
	}
	if c.Uni == nil {
		tracer.Panic(fmt.Errorf("%T.Uni must not be empty", c))
	}
	if c.Wrk == nil {
		tracer.Panic(fmt.Errorf("%T.Wrk must not be empty", c))
	}

	return &Engine{
		don: c.Don,
		fil: c.Fil,
		log: c.Log,
		rtr: c.Rtr,
		tkx: c.Tkx,
		uni: c.Uni,
		wrk: c.Wrk,

		lkp: newLookup(c.Cap),
		mem: newMemory(c.Cap),
		ply: newPlayer(c.Cap),
	}
}
