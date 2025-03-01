package client

import (
	"context"

	"github.com/coder/websocket"
)

// TODO:test sequence byte is added to any message

// Stream writes the given buffer to the underlying client connection using a
// sequence byte added at the end of the message. Clients can keep track of the
// sequence byte in order to ensure that they received all messages that the
// server intended to send to them.
func (c *Client) Stream(buf []byte) error {
	c.seq++
	return c.con.Write(context.Background(), websocket.MessageBinary, append(buf, c.seq))
}
