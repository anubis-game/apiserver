package client

import (
	"time"

	"github.com/coder/websocket"
	"github.com/ethereum/go-ethereum/common"
)

type Config struct {
	Con *websocket.Conn
	Wal common.Address
}

type Client struct {
	buf chan []byte
	exp chan struct{}
	rea chan struct{}
	tic chan struct{}
	wri chan struct{}

	cap int
	tiC <-chan time.Time

	con *websocket.Conn
	wal common.Address
}

func New(c Config) *Client {
	return &Client{
		buf: make(chan []byte, 1024), // closed never
		exp: make(chan struct{}),     // closed in server/handler/connect/client.go
		rea: make(chan struct{}),     // closed in server/handler/connect/client.go
		tic: make(chan struct{}),     // closed in client/daemon.go
		wri: make(chan struct{}),     // closed in client/daemon.go

		cap: 256,
		tiC: time.Tick(5 * time.Second),

		con: c.Con,
		wal: c.Wal,
	}
}
