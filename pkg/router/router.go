package router

import (
	"time"
)

type Router struct {
	cli *Client
	eng *Engine
}

func New() *Router {
	var cre chan Packet
	{
		cre = make(chan Packet, 500)
	}

	var del chan Packet
	{
		del = make(chan Packet, 500)
	}

	var fan <-chan time.Time
	{
		fan = time.NewTicker(25 * time.Millisecond).C
	}

	return &Router{
		cli: &Client{cre: cre, del: del},
		eng: &Engine{cre: cre, del: del, fan: fan},
	}
}

func (r *Router) Client() *Client {
	return r.cli
}

func (r *Router) Engine() *Engine {
	return r.eng
}
