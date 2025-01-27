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
	cre chan<- Packet
	del chan<- Packet
	lim *xsync.MapOf[common.Address, ratelimit.Limiter]
}

func (c *Client) Create(uid uuid.UUID, cli *client.Client) error {
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
		c.cre <- Packet{Byt: nil, Cli: cli, Uid: uid}
	}

	return nil
}

func (c *Client) Delete(uid uuid.UUID, cli *client.Client) {
	// Prevent DOS attacks and rate limit client specific stream input, so that
	// our internal fanout schedule cannot be overloaded maliciously.

	var lim ratelimit.Limiter
	{
		lim, _ = c.lim.LoadOrCompute(cli.Wallet(), newLim)
	}

	{
		lim.Take()
	}

	// Once a ticket was available for the client, we can proceed to enter the
	// synchronization loop.

	{
		c.del <- Packet{Byt: nil, Cli: cli, Uid: uid}
	}
}

func newLim() ratelimit.Limiter {
	return ratelimit.New(
		3,                                  // 1 move, 1 race, 1 buffer
		ratelimit.Per(25*time.Millisecond), // per standard frame
		ratelimit.WithSlack(0),             // without re-using unused capacity
	)
}
