package client

import (
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/xh3b4sd/tracer"
)

func (c *Client) ping(byt []byte) error {
	// We accept a single roundtrip byte that we echo back as is. If the input
	// buffer is not exactly of length 1, then we return an error and terminate
	// the client connection.

	if len(byt) != 1 {
		return tracer.Maskf(pingBufferInvalidError, "%d", len(byt))
	}

	// Encode the ping response and send the roundtrip byte back to the client.

	{
		c.fcn <- []byte{byte(schema.Pong), byt[0]}
	}

	return nil
}
