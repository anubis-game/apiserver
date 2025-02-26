package client

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/xh3b4sd/tracer"
)

const (
	// TTL is the maximum amount of time that every client is allowed to stay
	// connected in one session.
	TTL = 3 * time.Hour
)

func (c *Client) Daemon() {
	// Process every message that this client receives in its own goroutine. Any
	// write error causes the client connection to terminate. The buffer channel
	// defines a capacity of 1024 pending messages, but we terminate the client
	// connection already after 25% saturation. 40 messages may acumulate every
	// second at a frame rate of 25 ms/s. That means roughly 200 messages may
	// accumulate within a 5 second window, which means that we have at least 3
	// intervals to decide whether we want to terminate a congested client
	// connection.

	go func() {
		for {
			select {
			case <-c.don:
				return
			case <-c.rea:
				return
			case b := <-c.fcn:
				if c.logger(c.Stream(b)) != nil {
					close(c.wri)
					return
				}
			}
		}
	}()

	// Ensure the client connection gets terminated on excessive saturation.
	// Every 5 seconds we check for the current amount of buffer congestion. We
	// allow every client to accumulate 256 pending messages before we close the
	// ticker channel below, which then triggers the client termination in the
	// server handler.

	go func() {
		for {
			select {
			case <-c.don:
				return
			case <-c.rea:
				return
			case <-c.tiC:
				if len(c.fcn) >= c.cap {
					close(c.tic)
					return
				}
			}
		}
	}()

	// With setting up the client connection, we also setup a connection deadline
	// in order to limit the maximum amount of time that every client is allowed
	// to stay connected in one session.

	go func() {
		for {
			select {
			case <-c.don:
				return
			case <-c.rea:
				return
			case <-time.After(TTL):
				close(c.exp)
				return
			}
		}
	}()

	// The reader loop blocks this call until any error occurs. We try to log the
	// useful errors once the reader loop stops. At last we close the reader
	// channel, which is relevant for the cleanup of this client connection. See
	// also Engine.Delete().

	{
		c.logger(c.reader()) // nolint:errcheck
	}

	{
		close(c.rea)
	}
}

func (c *Client) logger(err error) error {
	if errors.Is(err, net.ErrClosed) {
		// fall through
	} else if err != nil {
		c.log.Log(
			"level", "error",
			"message", err.Error(),
			"stack", tracer.Stack(err),
		)
	}

	return err
}

func (c *Client) reader() error {
	for {
		// Prevent DOS attacks and rate limit client specific stream input, so that
		// our internal fanout schedule cannot be overloaded maliciously.

		{
			c.lim.Take()
		}

		// Read the next incoming message from the client connection.

		_, byt, err := c.con.Read(context.Background())
		if err != nil {
			return tracer.Mask(err)
		}

		// Dispatch according to the action byte.

		switch schema.Action(byt[0]) {
		case schema.Ping:
			err = c.ping(byt)
		case schema.Auth:
			err = c.auth(byt)
		case schema.Uuid:
			err = c.uuid(byt)
		case schema.Race:
			err = c.race(byt)
		case schema.Turn:
			err = c.turn(byt)
		}

		if err != nil {
			return tracer.Mask(err)
		}
	}
}
