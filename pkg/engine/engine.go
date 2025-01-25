package engine

import (
	"fmt"
	"runtime"
	"time"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/ethereum/go-ethereum/common"
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
	// cli contains all connected clients. This is a native Go map, and we
	// synchronize it via channel access in line with the time based fanout
	// procedure.
	cli map[common.Address]*client.Client
	// don is the global channel to signal program termination. If this channel is
	// closed, then all streaming connections should be terminated gracefully.
	don <-chan struct{}
	// log is a simple logger interface to print system relevant information.
	log logger.Interface
	// nrg contains the energy messages prepared to be sent out to all connected
	// clients during the time based fanout procedure.
	nrg *xsync.MapOf[common.Address, [][]byte]
	// ply contains the player messages prepared to be sent out to all connected
	// clients during the time based fanout procedure.
	ply *xsync.MapOf[common.Address, [][]byte]
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

	return &Engine{
		cli: map[common.Address]*client.Client{},
		don: c.Don,
		log: c.Log,
		nrg: xsync.NewMapOf[common.Address, [][]byte](),
		ply: xsync.NewMapOf[common.Address, [][]byte](),
		rtr: c.Rtr,
		sem: make(chan struct{}, runtime.NumCPU()),
	}
}
