package client

import (
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/xh3b4sd/tracer"
)

func (c *Client) turn(byt []byte) error {
	// If we do not receive exactly 3 bytes, then we simply return an error. The
	// three required bytes here are the action byte, the quadrant byte, and the
	// angle byte.

	if len(byt) != 3 {
		return tracer.Maskf(turnBytesInvalidError, "%d", len(byt))
	}

	// If the quadrant byte is not one of [1 2 3 4], then we simply ignore the
	// user input.

	if byt[1]-1 > 3 {
		return tracer.Maskf(turnQuadrantRangeError, "%#v", byt[1])
	}

	// Prevent DOS attacks and rate limit client specific stream input, so that
	// our internal fanout schedule cannot be overloaded maliciously.

	{
		c.lim.Take()
	}

	// Just send the turn signal to the engine for reconciliation.

	{
		c.rtr.Turn() <- router.Turn{Uid: c.uid, Qdr: byt[1], Agl: byt[2]}
	}

	return nil
}
