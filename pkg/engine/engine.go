package engine

import (
	"fmt"
	"time"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Cap int
	Don <-chan struct{}
	Fil *filler.Filler
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
	// fbf contains the fanout buffers ready to be sent out to every player during
	// the ticker based fanout procedure. Any respective byte slice may be empty,
	// or contain one, or multiple encoded messages.
	fbf *xsync.MapOf[byte, []byte]
	// fcn contains the fanout channels for every player. It is critically
	// important that modifications on fcn are only done sequentially by a single
	// writer.
	fcn []chan<- []byte
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
	// tkx
	tkx *tokenx.TokenX[common.Address]
	// tur
	tur []router.Turn
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
		fbf: xsync.NewMapOf[byte, []byte](),
		fcn: make([]chan<- []byte, c.Cap),
		fil: c.Fil,
		lkp: &lookup{
			nrg: xsync.NewMapOf[matrix.Partition, map[matrix.Coordinate]struct{}](),
			ply: xsync.NewMapOf[matrix.Partition, map[byte]struct{}](),
		},
		log: c.Log,
		mem: &memory{
			nrg: xsync.NewMapOf[matrix.Coordinate, *energy.Energy](),
			ply: xsync.NewMapOf[byte, *player.Player](),
		},
		rac: make([]byte, c.Cap),
		rtr: c.Rtr,
		tkx: c.Tkx,
		tur: make([]router.Turn, c.Cap),
		uni: c.Uni,
		wrk: c.Wrk,
	}
}
