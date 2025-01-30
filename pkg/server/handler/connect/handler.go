package connect

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/anubis-game/apiserver/pkg/cache"
	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

const (
	// max is the maximum amount of concurrent client connections accepted by the
	// stream engine.
	max = 500
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
}

type Handler struct {
	// ctx is the global context instance that we inject into every client struct.
	// We are not leveraging any of the underlying context specific control flow
	// primitives, but the websocket implementation that we are using requires a
	// context parameter to be provided.  And so in order to not garbage collect
	// useless context instances all the time, we define a single global context
	// and reuse that for the required websocket parameters everywhere.
	ctx context.Context
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

	var opt *websocket.AcceptOptions
	{
		opt = &websocket.AcceptOptions{
			InsecureSkipVerify: true, // TODO verify origin
			Subprotocols: []string{
				string(schema.DualHandshake),
				string(schema.UserChallenge),
			},
		}
	}

	return &Handler{
		ctx: context.Background(),
		don: c.Don,
		ind: cache.NewSxnc[common.Address, uuid.UUID](),
		log: c.Log,
		opt: opt,
		reg: c.Reg,
		rtr: c.Rtr,
		sem: make(chan struct{}, max),
		ttl: musDur(c.Env.ConnectionTimeout, "s"),
		txp: cache.NewTime[uuid.UUID](),
		tok: cache.NewSxnc[uuid.UUID, common.Address](),
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
