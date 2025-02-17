package router

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/ethereum/go-ethereum/common"
	"github.com/puzpuzpuz/xsync/v3"
	"go.uber.org/ratelimit"
)

type Router struct {
	cli *Client
	eng *Engine
}

// Router is the bridge between server endpoint and game engine, allowing us to
// separate client connections and game state. Note that the *time.Ticker used
// for our fanout procedure is never stopped, because this ticker is used across
// the lifetime of the entire game engine.
func New() *Router {
	var joi chan Packet
	{
		joi = make(chan Packet, 500)
	}

	var mov chan Packet
	{
		mov = make(chan Packet, 2000)
	}

	var rac chan Packet
	{
		rac = make(chan Packet, 2000)
	}

	var psh <-chan time.Time
	{
		psh = time.NewTicker(vector.Frm * time.Millisecond).C
	}

	var lim *xsync.MapOf[common.Address, ratelimit.Limiter]
	{
		lim = xsync.NewMapOf[common.Address, ratelimit.Limiter]()
	}

	return &Router{
		cli: &Client{joi: joi, mov: mov, rac: rac, lim: lim},
		eng: &Engine{joi: joi, mov: mov, rac: rac, psh: psh},
	}
}

func (r *Router) Client() *Client {
	return r.cli
}

func (r *Router) Engine() *Engine {
	return r.eng
}
