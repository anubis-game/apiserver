package stream

import (
	"context"
	"fmt"
	"time"

	"github.com/anubis-game/apiserver/pkg/cache"
	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/coder/websocket"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Don is the global channel to signal program termination. If this channel is
	// closed, then all streaming connections should be terminated gracefully.
	Don <-chan struct{}
	// Log is the logger interface for printing structured log messages.
	Log logger.Interface
	// Out is the connection timeout that the stream engine should enforce upon
	// connected clients. All associated onchain and offchain resources must be
	// released after having served clients successfully for this amount of time.
	Out time.Duration
	// Reg is the onchain interface for the Registry smart contract.
	Reg *registry.Registry
}

type Stream struct {
	cli cache.Interface[Client]
	ctx context.Context
	don <-chan struct{}
	exp *cache.Time
	ind cache.Interface[string]
	log logger.Interface
	opt *websocket.AcceptOptions
	reg *registry.Registry
	tok cache.Interface[string]
}

func New(c Config) *Stream {
	if c.Don == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Don must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Out == 0 {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Out must not be empty", c)))
	}
	if c.Reg == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Reg must not be empty", c)))
	}

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var exp *cache.Time
	{
		exp = cache.NewTime(c.Out)
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

	{
		go exp.Expire(time.Minute)
	}

	return &Stream{
		cli: cache.NewData[Client](),
		ctx: ctx,
		don: c.Don,
		exp: exp,
		ind: cache.NewData[string](),
		log: c.Log,
		opt: opt,
		reg: c.Reg,
		tok: cache.NewData[string](),
	}
}
