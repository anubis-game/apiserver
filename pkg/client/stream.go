package client

import "github.com/coder/websocket"

func (c *Client) Stream(byt []byte) {
	err := c.con.Write(c.ctx, websocket.MessageBinary, byt)
	if err != nil {
		close(c.Writer())
	}
}
