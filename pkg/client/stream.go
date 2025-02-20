package client

import (
	"context"
	"time"

	"github.com/coder/websocket"
	"github.com/xh3b4sd/tracer"
)

const (
	// Max is the maximum duration that any websocket write is allowed to take.
	// Any client connection blocking writes longer than this amount of time will
	// be forcibly closed.
	Max = 5 * time.Second
)

func (c *Client) Stream(byt []byte) error {
	// Setup the prerequisites for a timeout limited write stream. It doesn't
	// matter where Client.Stream originates from, we always have to ensure that
	// no single client write blocks the global runtime.

	var ctx context.Context
	var don context.CancelFunc
	{
		ctx, don = context.WithTimeout(context.Background(), Max)
	}

	// Write the given binary message to this client connection.

	err := c.con.Write(ctx, websocket.MessageBinary, byt)
	if err != nil {
		don()
		return tracer.Mask(err)
	}

	// In case the client write above finished early, cleanup the context related
	// resources incurred during context creation.

	{
		don()
	}

	return nil
}
