package client

import (
	"github.com/anubis-game/apiserver/pkg/schema"
	"github.com/google/uuid"
	"github.com/xh3b4sd/tracer"
)

func (c *Client) auth(_ []byte) error {
	var err error

	// Prevent DOS attacks and rate limit client specific stream input, so that
	// our internal fanout schedule cannot be overloaded maliciously.

	{
		c.lim.Take()
	}

	// Create a new session token using V4 UUIDs for the requesting Wallet
	// address. If the requesting wallet maps to an existing session token, then
	// the existing token will be deleted, meaning it cannot be used anymore.
	// TokenX.Create() always issues new tokens.

	var tok uuid.UUID
	{
		tok, err = c.tkx.Create(c.wal)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	// Encode the auth response and send the new session token back to the client
	// connection that requested this new credential.

	{
		c.fcn <- schema.Encode(schema.Auth, tok[:])
	}

	return nil
}
