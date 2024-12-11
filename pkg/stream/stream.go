package stream

import (
	"context"
	"fmt"
	"sync"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/coder/websocket"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don chan struct{}
	Log logger.Interface
	Reg *registry.Registry
}

type Stream struct {
	cli map[string]Client
	ctx context.Context
	don chan struct{}
	log logger.Interface
	mut sync.RWMutex
	opt *websocket.AcceptOptions
	reg *registry.Registry
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

	var ctx context.Context
	{
		ctx = context.Background()
	}

	var opt *websocket.AcceptOptions
	{
		opt = &websocket.AcceptOptions{
			InsecureSkipVerify: true, // TODO verify origin
			Subprotocols:       []string{"dual-handshake"},
		}
	}

	return &Stream{
		cli: map[string]Client{},
		ctx: ctx,
		don: c.Don,
		log: c.Log,
		mut: sync.RWMutex{},
		opt: opt,
		reg: c.Reg,
	}
}
