package client

import (
	"context"

	"github.com/anubis-game/apiserver/pkg/window"
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/ratelimit"
)

type Config struct {
	Con *websocket.Conn
	Ctx context.Context
	Lim ratelimit.Limiter
	Wal common.Address
	Win *window.Window
}

type Client struct {
	exp chan struct{}
	rea chan struct{}
	wri chan struct{}

	con *websocket.Conn
	ctx context.Context
	lim ratelimit.Limiter
	wal common.Address
	win *window.Window
}

func New(c Config) *Client {
	return &Client{
		exp: make(chan struct{}),
		rea: make(chan struct{}),
		wri: make(chan struct{}),

		con: c.Con,
		ctx: c.Ctx,
		lim: c.Lim,
		wal: c.Wal,
		win: c.Win,
	}
}

func (c *Client) Expiry() chan struct{} {
	return c.exp
}

func (c *Client) Reader() chan struct{} {
	return c.rea
}

func (c *Client) Stream(byt []byte) {
	err := c.con.Write(c.ctx, websocket.MessageBinary, byt)
	if err != nil {
		close(c.Writer())
	}
}

func (c *Client) Ticket() {
	c.lim.Take()
}

func (c *Client) Wallet() common.Address {
	return c.wal
}

func (c *Client) Window() *window.Window {
	return c.win
}

func (c *Client) Writer() chan struct{} {
	return c.wri
}
