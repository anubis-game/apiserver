package connect

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/random"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Agl
	Agl *random.Random
	// Cap
	Cap int
	// Crd
	Crd *random.Random
	// Don is the global channel to signal program termination. If this channel is
	// closed, then all streaming connections should be terminated gracefully.
	Don <-chan struct{}
	// Log is the logger interface for printing structured log messages.
	Log logger.Interface
	// Qdr
	Qdr *random.Random
	// Reg is the onchain interface for the Registry smart contract.
	Reg *registry.Registry
	// Rtr
	Rtr *router.Client
	// Tkx
	Tkx *tokenx.TokenX[common.Address]
	// Uni provides a thread safe mechanism to allocate compact player IDs.
	// Allocation happens in the server handler, freeing allocated player IDs
	// happens in the game engine.
	Uni *unique.Unique[common.Address, byte]
}

type Handler struct {
	agl *random.Random
	crd *random.Random
	don <-chan struct{}
	log logger.Interface
	opt *websocket.AcceptOptions
	qdr *random.Random
	reg *registry.Registry
	rtr *router.Client
	sem chan struct{}
	tkx *tokenx.TokenX[common.Address]
	uni *unique.Unique[common.Address, byte]
}

func New(c Config) *Handler {
	if c.Agl == nil {
		tracer.Panic(fmt.Errorf("%T.Agl must not be empty", c))
	}
	if c.Crd == nil {
		tracer.Panic(fmt.Errorf("%T.Crd must not be empty", c))
	}
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}
	if c.Qdr == nil {
		tracer.Panic(fmt.Errorf("%T.Qdr must not be empty", c))
	}
	if c.Reg == nil {
		tracer.Panic(fmt.Errorf("%T.Reg must not be empty", c))
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

	var sem chan struct{}
	{
		sem = make(chan struct{}, c.Cap)
	}

	return &Handler{
		agl: c.Agl,
		crd: c.Crd,
		don: c.Don,
		log: c.Log,
		qdr: c.Qdr,
		reg: c.Reg,
		rtr: c.Rtr,
		tkx: c.Tkx,
		uni: c.Uni,

		opt: opt,
		sem: sem,
	}
}
