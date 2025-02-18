package client

import (
	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Con *websocket.Conn
	Wal common.Address
}

type Client struct {
	exp chan struct{}
	rea chan struct{}
	wri chan struct{}

	con *websocket.Conn
	wal common.Address
}

func New(c Config) *Client {
	return &Client{
		exp: make(chan struct{}),
		rea: make(chan struct{}),
		wri: make(chan struct{}),

		con: c.Con,
		wal: c.Wal,
	}
}
