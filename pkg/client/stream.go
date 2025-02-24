package client

import (
	"context"

	"github.com/coder/websocket"
)

func (c *Client) Stream(byt []byte) error {
	return c.con.Write(context.Background(), websocket.MessageBinary, byt)
}
