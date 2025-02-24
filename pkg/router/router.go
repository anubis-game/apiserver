package router

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/vector"
	"github.com/ethereum/go-ethereum/common"
	"github.com/puzpuzpuz/xsync/v3"
	"go.uber.org/ratelimit"
)

type Config struct {
	Env envvar.Env
}

type Router struct {
	cli *Client
	eng *Engine
}

// Router is the bridge between server endpoint and game engine, allowing us to
// separate client connections and game state. Note that the *time.Ticker used
// for our fanout procedure is never stopped, because this ticker is used across
// the lifetime of the entire game engine.
func New(c Config) *Router {
	var uid chan Packet
	{
		uid = make(chan Packet, c.Env.EngineCapacity)
	}

	var rac chan Packet
	{
		rac = make(chan Packet, c.Env.EngineCapacity*2)
	}

	var tur chan Packet
	{
		tur = make(chan Packet, c.Env.EngineCapacity*2)
	}

	var tic <-chan time.Time
	{
		tic = time.NewTicker(vector.Frm * time.Millisecond).C
	}

	var lim *xsync.MapOf[common.Address, ratelimit.Limiter]
	{
		lim = xsync.NewMapOf[common.Address, ratelimit.Limiter]()
	}

	return &Router{
		cli: &Client{uid: uid, rac: rac, tur: tur, lim: lim},
		eng: &Engine{uid: uid, rac: rac, tur: tur, tic: tic},
	}
}

func (r *Router) Client() *Client {
	return r.cli
}

func (r *Router) Engine() *Engine {
	return r.eng
}
