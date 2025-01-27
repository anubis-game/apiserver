package engine

import (
	"fmt"
	"runtime"
	"time"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/player"
	"github.com/anubis-game/apiserver/pkg/random"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/google/uuid"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don <-chan struct{}
	Log logger.Interface
	Rtr *router.Engine
}

type Engine struct {
	// buffer
	buf *buffer
	// don is the global channel to signal program termination. If this channel is
	// closed, then all streaming connections should be terminated gracefully.
	don <-chan struct{}
	// filler
	fil *filler
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
	// tic is the global fanout ticker. The first tick is initialized in Daemon().
	tic time.Time
}

func New(c Config) *Engine {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}

	var ang *random.Random
	{
		ang = random.New(random.Config{
			Buf: 500,
			Don: c.Don,
			Log: c.Log,
			Max: 255,
			Min: 0,
		})
	}

	var crd *random.Random
	{
		crd = random.New(random.Config{
			Buf: 3000,
			Don: c.Don,
			Log: c.Log,
			Max: matrix.Max,
			Min: matrix.Min,
		})
	}

	var qdr *random.Random
	{
		qdr = random.New(random.Config{
			Buf: 500,
			Don: c.Don,
			Log: c.Log,
			Max: 4,
			Min: 1,
		})
	}

	return &Engine{
		buf: &buffer{
			nrg: xsync.NewMapOf[uuid.UUID, [][]byte](),
			ply: xsync.NewMapOf[uuid.UUID, [][]byte](),
		},
		don: c.Don,
		fil: &filler{
			ang: ang,
			crd: crd,
			qdr: qdr,
		},
		lkp: &lookup{
			nrg: xsync.NewMapOf[matrix.Bucket, uuid.UUID](),
			ply: xsync.NewMapOf[matrix.Bucket, uuid.UUID](),
		},
		log: c.Log,
		mem: &memory{
			nrg: map[uuid.UUID]*energy.Energy{},
			ply: map[uuid.UUID]*player.Player{},
		},
		rtr: c.Rtr,
		sem: make(chan struct{}, runtime.NumCPU()),
	}
}
