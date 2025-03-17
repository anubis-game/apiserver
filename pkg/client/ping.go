package client

import (
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/xh3b4sd/tracer"
)

func (c *Client) ping(byt []byte) error {
	// We accept a single roundtrip byte that we echo back as is. The two required
	// bytes here are the action byte, and the roundtrip byte.

	if len(byt) != 2 {
		return tracer.Maskf(pingBufferInvalidError, "%d", len(byt))
	}

	// Encode the ping response and send the roundtrip byte back to the client.

	{
		c.fcn <- []byte{byte(schema.Pong), byt[1]}
	}

	return nil
}
