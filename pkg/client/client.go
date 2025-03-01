package client

import (
	"fmt"
	"time"

	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
	"go.uber.org/ratelimit"
)

type Config struct {
	Con *websocket.Conn
	Don <-chan struct{}
	Fcn chan []byte
	Lim ratelimit.Limiter
	Log logger.Interface
	Rtr *router.Client
	Tkx *tokenx.TokenX[common.Address]
	Uid byte
	Wal common.Address
}

type Client struct {
	exp chan struct{}
	rea chan struct{}
	tic chan struct{}
	wri chan struct{}

	cap int
	seq byte
	tiC <-chan time.Time

	con *websocket.Conn
	don <-chan struct{}
	fcn chan []byte
	lim ratelimit.Limiter
	log logger.Interface
	rtr *router.Client
	tkx *tokenx.TokenX[common.Address]
	uid byte
	wal common.Address
}

func New(c Config) *Client {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}
	if c.Tkx == nil {
		tracer.Panic(fmt.Errorf("%T.Tkx must not be empty", c))
	}

	return &Client{
		exp: make(chan struct{}), // closed in client/daemon.go
		rea: make(chan struct{}), // closed in client/daemon.go
		tic: make(chan struct{}), // closed in client/daemon.go
		wri: make(chan struct{}), // closed in client/daemon.go

		cap: 256,
		seq: 0,
		tiC: time.Tick(5 * time.Second),

		con: c.Con,
		don: c.Don,
		fcn: c.Fcn,
		lim: c.Lim,
		log: c.Log,
		rtr: c.Rtr,
		tkx: c.Tkx,
		uid: c.Uid,
		wal: c.Wal,
	}
}
