package client

import (
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/xh3b4sd/tracer"
)

func (c *Client) uuid(byt []byte) error {
	// If we do not receive exactly 1 byte, then we simply return an error. The
	// one required byte here is only the action byte.

	if len(byt) != 1 {
		return tracer.Maskf(uuidBytesInvalidError, "%d", len(byt))
	}

	// Just send the turn signal to the engine for reconciliation.

	// TODO:test unit test the garbage collection of leaving players using the
	// router.Drop byte.

	{
		c.rtr.Uuid() <- router.Uuid{Uid: c.uid, Jod: router.Join, Wal: c.wal, Fcn: c.fcn}
	}

	return nil
}
