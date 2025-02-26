package router

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/vector"
)

type Config struct {
	Env envvar.Env
}

type Router struct {
	cli *Client
	eng *Engine
}

// Router is the bridge between server endpoint and game engine, allowing us to
// separate client connections and game state.
func New(c Config) *Router {
	var rac chan byte
	var tur chan Turn
	var uid chan Uuid
	{
		rac = make(chan byte, c.Env.EngineCapacity*2)
		tur = make(chan Turn, c.Env.EngineCapacity*2)
		uid = make(chan Uuid, c.Env.EngineCapacity)
	}

	var tic <-chan time.Time
	{
		tic = time.Tick(vector.Frm * time.Millisecond)
	}

	return &Router{
		cli: &Client{uid: uid, rac: rac, tur: tur},
		eng: &Engine{uid: uid, rac: rac, tur: tur, tic: tic},
	}
}

func (r *Router) Client() *Client {
	return r.cli
}

func (r *Router) Engine() *Engine {
	return r.eng
}
