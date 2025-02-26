package client

import (
	"github.com/xh3b4sd/tracer"
)

func (c *Client) race(byt []byte) error {
	// If we do not receive exactly 1 byte, then we simply return an error. The
	// one required byte here is only the action byte.

	if len(byt) != 1 {
		return tracer.Maskf(raceBytesInvalidError, "%d", len(byt))
	}

	// Just send the race signal to the engine for reconciliation.

	{
		c.rtr.Race() <- c.uid
	}

	return nil
}
