package client

import (
	"context"

	"github.com/coder/websocket"
)

type Config struct {
	Con *websocket.Conn
	Ctx context.Context
}

type Client struct {
	exp chan struct{}
	rea chan struct{}
	wri chan struct{}

	con *websocket.Conn
	ctx context.Context
}

func New(c Config) *Client {
	return &Client{
		exp: make(chan struct{}),
		rea: make(chan struct{}),
		wri: make(chan struct{}),

		con: c.Con,
		ctx: c.Ctx,
	}
}

func (c *Client) Expiry() chan struct{} {
	return c.exp
}

func (c *Client) Reader() chan struct{} {
	return c.rea
}

func (c *Client) Writer() chan struct{} {
	return c.wri
}

func (c *Client) Stream(byt []byte) {
	err := c.con.Write(c.ctx, websocket.MessageBinary, byt)
	if err != nil {
		close(c.Writer())
	}
}
