package engine

import (
	"fmt"
	"runtime"
	"time"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don <-chan struct{}
	Fil *filler.Filler
	Log logger.Interface
	Rtr *router.Engine
	Wrk worker.Ensure
}

type Engine struct {
	// buffer
	buf *buffer
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
	// rtr is the bridge synchronizing the server handler and the game engine
	rtr *router.Engine
	// sem is the global worker ticket. The capacity of this semaphore channel
	// limits the amount of work we can do concurrently.
	sem chan struct{}
	// tic is the global pointer keeping track of the fanout related time ticks.
	// This timestamp tracks at which point the latest fanout procedure has been
	// executed. The first tick is initialized in Daemon().
	tic time.Time
	// tim is the amount of time that a single cycle of client streams is allowed
	// to take. Any client taking longer to receive data packets than this
	// deadline specifies will stop to be served for the blocked cycle. Note that
	// single writes are globally limited to 5 seconds of allowed write time. So
	// any client connection that blocks writes for more than 5 seconds per
	// message will be terminated, which also means that those clients will be
	// removed from the game that they are playing.
	tim time.Duration
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
	if c.Wrk == nil {
		tracer.Panic(fmt.Errorf("%T.Wrk must not be empty", c))
	}

	return &Engine{
		buf: &buffer{
			nrg: xsync.NewMapOf[[2]byte, [][]byte](),
			ply: xsync.NewMapOf[[2]byte, [][]byte](),
		},
		don: c.Don,
		fil: c.Fil,
		lkp: &lookup{
			nrg: xsync.NewMapOf[object.Object, map[object.Object]struct{}](),
			ply: xsync.NewMapOf[object.Object, map[[2]byte]struct{}](),
		},
		log: c.Log,
		mem: &memory{
			nrg: map[object.Object]*energy.Energy{},
			ply: map[[2]byte]*player.Player{},
		},
		rtr: c.Rtr,
		sem: make(chan struct{}, runtime.NumCPU()),
		wrk: c.Wrk,
	}
}
