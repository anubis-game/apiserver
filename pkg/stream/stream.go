package stream

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/anubis-game/apiserver/pkg/cache"
	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/anubis-game/apiserver/pkg/worker/release"
	"github.com/anubis-game/apiserver/pkg/worker/resolve"
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
	//
	Env envvar.Env
	// Log is the logger interface for printing structured log messages.
	Log logger.Interface
	// Reg is the onchain interface for the Registry smart contract.
	Reg *registry.Registry
	//
	Rel worker.Create[release.Packet]
	//
	Res worker.Create[resolve.Packet]
}

type Stream struct {
	cli cache.Interface[common.Address, *client.Client]
	ctx context.Context
	don <-chan struct{}
	ind cache.Interface[common.Address, uuid.UUID]
	log logger.Interface
	opt *websocket.AcceptOptions
	reg *registry.Registry
	rel worker.Create[release.Packet]
	res worker.Create[resolve.Packet]
	sem chan struct{}
	// ttl is the connection timeout that the stream engine should enforce upon
	// connected clients. All associated onchain and offchain resources must be
	// released after having served clients successfully for this amount of time.
	ttl time.Duration
	txp *cache.Time[uuid.UUID]
	tok cache.Interface[uuid.UUID, common.Address]
	wxp *cache.Time[common.Address]
}

func New(c Config) *Stream {
	if c.Don == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Don must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Reg == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Reg must not be empty", c)))
	}
	if c.Rel == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rel must not be empty", c)))
	}
	if c.Res == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Res must not be empty", c)))
	}

	var ctx context.Context
	{
		ctx = context.Background()
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

	return &Stream{
		cli: cache.NewPool[common.Address, *client.Client](),
		ctx: ctx,
		don: c.Don,
		ind: cache.NewSxnc[common.Address, uuid.UUID](),
		log: c.Log,
		opt: opt,
		reg: c.Reg,
		rel: c.Rel,
		res: c.Res,
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
