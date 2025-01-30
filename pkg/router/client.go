package router

import (
	"fmt"
	"time"

	"github.com/anubis-game/apiserver/pkg/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/puzpuzpuz/xsync/v3"
	"github.com/xh3b4sd/tracer"
	"go.uber.org/ratelimit"
)

type Client struct {
	joi chan<- Packet
	mov chan<- Packet
	rac chan<- Packet

	lim *xsync.MapOf[common.Address, ratelimit.Limiter]
}

func (c *Client) Join(uid uuid.UUID, cli *client.Client, _ []byte) error {
	// Prevent DOS attacks and rate limit client specific stream input, so that
	// our internal fanout schedule cannot be overloaded maliciously.

	var lim ratelimit.Limiter
	var exi bool
	{
		lim, exi = c.lim.LoadOrCompute(cli.Wallet(), newLim)
	}

	if exi {
		return tracer.Mask(fmt.Errorf("client %q already joined", cli.Wallet()))
	}

	{
		lim.Take()
	}

	// Once a ticket was available for the client, we can proceed to enter the
	// synchronization loop.

	{
		c.joi <- Packet{Byt: nil, Cli: cli, Uid: uid}
	}

	return nil
}

func (c *Client) Move(uid uuid.UUID, cli *client.Client, byt []byte) error {
	// Prevent DOS attacks and rate limit client specific stream input, so that
	// our internal fanout schedule cannot be overloaded maliciously.

	var lim ratelimit.Limiter
	var exi bool
	{
		lim, exi = c.lim.LoadOrCompute(cli.Wallet(), newLim)
	}

	if !exi {
		return tracer.Mask(fmt.Errorf("client %q not joined", cli.Wallet()))
	}

	{
		lim.Take()
	}

	// Once a ticket was available for the client, we can proceed to enter the
	// synchronization loop.

	{
		c.mov <- Packet{Byt: byt, Cli: nil, Uid: uid}
	}

	return nil
}

func (c *Client) Race(uid uuid.UUID, cli *client.Client, _ []byte) error {
	// Prevent DOS attacks and rate limit client specific stream input, so that
	// our internal fanout schedule cannot be overloaded maliciously.

	var lim ratelimit.Limiter
	var exi bool
	{
		lim, exi = c.lim.LoadOrCompute(cli.Wallet(), newLim)
	}

	if !exi {
		return tracer.Mask(fmt.Errorf("client %q not joined", cli.Wallet()))
	}

	{
		lim.Take()
	}

	// Once a ticket was available for the client, we can proceed to enter the
	// synchronization loop.

	{
		c.rac <- Packet{Byt: nil, Cli: nil, Uid: uid}
	}

	return nil
}

func newLim() ratelimit.Limiter {
	return ratelimit.New(
		3,                                  // 1 move, 1 race, 1 buffer
		ratelimit.Per(25*time.Millisecond), // per standard frame
		ratelimit.WithSlack(0),             // without re-using unused capacity
	)
}
