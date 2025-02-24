package engine

import (
	"fmt"
	"time"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don <-chan struct{}
	Env envvar.Env
	Fil *filler.Filler
	Log logger.Interface
	Rtr *router.Engine
	Uni *unique.Unique[common.Address, byte]
	Wrk worker.Ensure
}

type Engine struct {
	// buf contains various messages prepared to be sent out to this player's
	// client during the time based fanout procedure. The byte slice may contain
	// multiple encoded messages.
	buf *xsync.MapOf[byte, []byte]
	// don is the global channel to signal program termination. If this channel is
	// closed, then all streaming connections should be terminated gracefully.
	don <-chan struct{}
	// filler
	fil *filler.Filler
	// lkp
	lkp *lookup
	// log is a simple logger interface to print system relevant information.
	log logger.Interface
	// mem
	mem *memory
	// rac
	rac []byte
	// rtr is the bridge synchronizing the server handler and the game engine
	rtr *router.Engine
	// tic is the global pointer keeping track of the fanout related time ticks.
	// This timestamp tracks at which point the latest fanout procedure has been
	// executed. The first tick is initialized in Engine.Daemon().
	tic time.Time
	// tur
	tur []Turn
	// uni
	uni *unique.Unique[common.Address, byte]
	// wrk
	wrk worker.Ensure
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
	if c.Uni == nil {
		tracer.Panic(fmt.Errorf("%T.Uni must not be empty", c))
	}
	if c.Wrk == nil {
		tracer.Panic(fmt.Errorf("%T.Wrk must not be empty", c))
	}

	return &Engine{
		buf: xsync.NewMapOf[byte, []byte](),
		don: c.Don,
		fil: c.Fil,
		lkp: &lookup{
			nrg: xsync.NewMapOf[object.Object, map[object.Object]struct{}](),
			ply: xsync.NewMapOf[object.Object, map[byte]struct{}](),
		},
		log: c.Log,
		mem: &memory{
			nrg: xsync.NewMapOf[object.Object, *energy.Energy](),
			ply: xsync.NewMapOf[byte, *player.Player](),
		},
		rac: make([]byte, c.Env.EngineCapacity),
		rtr: c.Rtr,
		tur: make([]Turn, c.Env.EngineCapacity),
		uni: c.Uni,
		wrk: c.Wrk,
	}
}
