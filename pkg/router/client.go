package router

import (
	"github.com/anubis-game/apiserver/pkg/client"
)

type Client struct {
	cre chan<- Packet
	del chan<- Packet
}

func (c *Client) Create(cli *client.Client) {
	// Prevent DOS attacks and rate limit client specific stream input, so that
	// our internal fanout schedule cannot be overloaded maliciously.

	{
		cli.Ticket()
	}

	// Once a ticket was available for the client, we can proceed to enter the
	// synchronization loop.

	{
		c.cre <- Packet{Byt: nil, Cli: cli}
	}
}

func (c *Client) Delete(cli *client.Client) {
	// Prevent DOS attacks and rate limit client specific stream input, so that
	// our internal fanout schedule cannot be overloaded maliciously.

	{
		cli.Ticket()
	}

	// Once a ticket was available for the client, we can proceed to enter the
	// synchronization loop.

	{
		c.del <- Packet{Byt: nil, Cli: cli}
	}
}
