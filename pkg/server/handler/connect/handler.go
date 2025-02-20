package connect

import (
	"fmt"
	"strings"
	"time"

	"github.com/anubis-game/apiserver/pkg/cache"
	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Don is the global channel to signal program termination. If this channel is
	// closed, then all streaming connections should be terminated gracefully.
	Don <-chan struct{}
	// Env
	Env envvar.Env
	// Log is the logger interface for printing structured log messages.
	Log logger.Interface
	// Reg is the onchain interface for the Registry smart contract.
	Reg *registry.Registry
	// Rtr
	Rtr *router.Client
	// Uni provides a thread safe mechanism to allocate compact player IDs.
	// Allocation happens in the server handler, freeing allocated player IDs
	// happens in the game engine.
	Uni *unique.Unique[common.Address]
}

type Handler struct {
	don <-chan struct{}
	ind cache.Interface[common.Address, uuid.UUID]
	log logger.Interface
	opt *websocket.AcceptOptions
	reg *registry.Registry
	// rtr is the bridge synchronizing the server handler and the game engine
	rtr *router.Client
	sem chan struct{}
	// ttl is the connection timeout that the stream engine should enforce upon
	// connected clients. All associated onchain and offchain resources must be
	// released after having served clients successfully for this amount of time.
	ttl time.Duration
	txp *cache.Time[uuid.UUID]
	tok cache.Interface[uuid.UUID, common.Address]
	uni *unique.Unique[common.Address]
	wxp *cache.Time[common.Address]
}

func New(c Config) *Handler {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}
	if c.Reg == nil {
		tracer.Panic(fmt.Errorf("%T.Reg must not be empty", c))
	}
	if c.Rtr == nil {
		tracer.Panic(fmt.Errorf("%T.Rtr must not be empty", c))
	}
	if c.Uni == nil {
		tracer.Panic(fmt.Errorf("%T.Uni must not be empty", c))
	}

	var opt *websocket.AcceptOptions
	{
		opt = &websocket.AcceptOptions{
			InsecureSkipVerify: true, // TODO:prod verify request origin in production
			Subprotocols: []string{
				string(schema.DualHandshake),
				string(schema.UserChallenge),
			},
		}
	}

	return &Handler{
		don: c.Don,
		ind: cache.NewSxnc[common.Address, uuid.UUID](),
		log: c.Log,
		opt: opt,
		reg: c.Reg,
		rtr: c.Rtr,
		sem: make(chan struct{}, c.Env.EngineCapacity),
		ttl: musDur(c.Env.ConnectionTimeout, "s"),
		txp: cache.NewTime[uuid.UUID](),
		tok: cache.NewSxnc[uuid.UUID, common.Address](),
		uni: c.Uni,
		wxp: cache.NewTime[common.Address](),
	}
}

func musDur(str string, uni string) time.Duration {
	dur, err := time.ParseDuration(strings.ReplaceAll(str, "_", "") + uni)
	if err != nil {
		tracer.Panic(err)
	}

	return dur
}
