package engine

import (
	"fmt"
	"time"

	"github.com/anubis-game/apiserver/pkg/energy"
	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/object"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/screen"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/ethereum/go-ethereum/common"
	"github.com/puzpuzpuz/xsync/v3"
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

// TODO:infra since we replaced a lot of synchronized maps with unsynchronized
// slices, we probably have to turn the slice values into atomic pointers.
// Testing concurrency in those relevant areas should yield a lot of data races,
// given the current code.

type Engine struct {
	// don is the global channel to signal program termination. If this channel is
	// closed, then all streaming connections should be terminated gracefully.
	don <-chan struct{}
	// fbf contains the fanout buffers ready to be sent out to every player during
	// the ticker based fanout procedure. Any respective byte slice may be empty,
	// or contain one, or multiple encoded messages.
	fbf [][]byte
	// fcn contains the fanout channels for every player. It is critically
	// important that modifications on fcn are only done sequentially by a single
	// writer.
	fcn []chan<- []byte
	// filler
	fil filler.Interface
	// fuw contains the cached fanout buffers mapping the byte ID and wallet
	// address for every active player. Using this lookup table we do not have to
	// do the same computation all over again every time.
	fuw [][]byte
	// log is a simple logger interface to print system relevant information.
	log logger.Interface
	// men contains all active energy packages currently placed within the game
	// map. Energy is identified by its precise X and Y coordinates, because only
	// one energy package can be in the same place at the same time. We can refer
	// to energy packages using their position only, because energy packages don't
	// move.
	men *xsync.MapOf[object.Object, *energy.Energy]
	// mvc contains the Vectors of all active players, indexed by their respective
	// byte IDs.
	mvc []*vector.Vector
	// pen contains the locations of energy packages indexed by partition
	// coordinates. Using this lookup table we can find all energy packages within
	// the given partition.
	//
	//     key:      partition coordinate
	//     value:    energy location
	//
	pen *xsync.MapOf[object.Object, map[object.Object]struct{}]
	// pvc contains the locations of player vectors indexed by partition
	// coordinates. Using this lookup table we can find all player vectors within
	// the given partition.
	//
	//     key:      partition coordinate
	//     value:    byte ID
	//
	pvc map[object.Object]map[byte]struct{}
	// rac
	rac []byte
	// rtr is the bridge synchronizing the server handler and the game engine
	rtr *router.Engine
	// scr
	scr []*screen.Screen
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
		fbf: make([][]byte, c.Cap),
		fcn: make([]chan<- []byte, c.Cap),
		fil: c.Fil,
		fuw: make([][]byte, c.Cap),
		log: c.Log,
		men: xsync.NewMapOf[object.Object, *energy.Energy](),
		mvc: make([]*vector.Vector, c.Cap),
		pen: xsync.NewMapOf[object.Object, map[object.Object]struct{}](),
		pvc: map[object.Object]map[byte]struct{}{},
		rac: make([]byte, c.Cap),
		rtr: c.Rtr,
		tkx: c.Tkx,
		tur: make([]router.Turn, c.Cap),
		uni: c.Uni,
		wrk: c.Wrk,
	}
}
