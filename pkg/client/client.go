package client

import (
	"context"

	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Con *websocket.Conn
	Ctx context.Context
	Wal common.Address
}

type Client struct {
	exp chan struct{}
	rea chan struct{}
	wri chan struct{}

	con *websocket.Conn
	ctx context.Context
	wal common.Address
}

func New(c Config) *Client {
	return &Client{
		exp: make(chan struct{}),
		rea: make(chan struct{}),
		wri: make(chan struct{}),

		con: c.Con,
		ctx: c.Ctx,
		wal: c.Wal,
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

func (c *Client) Wallet() common.Address {
	return c.wal
}

func (c *Client) Writer() chan struct{} {
	return c.wri
}
