package client

import (
	"errors"
	"net"

	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/xh3b4sd/tracer"
)

func (c *Client) Delete() {
	// Try to close the client connection, if it is not already closed, or
	// closing.

	err := c.con.CloseNow()
	if errors.Is(err, net.ErrClosed) {
		// fall through
	} else if err != nil {
		c.log.Log(
			"level", "error",
			"message", err.Error(),
			"stack", tracer.Stack(err),
		)
	}

	// After the client connection cannot accept any more incoming messages, we
	// wait for our internal reader channel to close too. Once the reader channel
	// is closed, we know for certain that no more messages can be forwarded to
	// the game engine.

	{
		<-c.rea
	}

	// Once no more external messages can be forwarded, send a drop signal to the
	// game engine so that this terminating client can be marked offline. Should
	// the player still be part of the game, then the game continues for that
	// player without interruption until the player dies. The player may establish
	// a new connection quickly and resume control of their active character.

	{
		c.rtr.Uuid() <- router.Uuid{Uid: c.uid, Jod: router.Drop}
	}
}
