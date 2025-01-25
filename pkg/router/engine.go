package router

import (
	"time"
)

type Engine struct {
	cre <-chan Packet
	del <-chan Packet
	fan <-chan time.Time
}

func (e *Engine) Create() <-chan Packet {
	return e.cre
}

func (e *Engine) Delete() <-chan Packet {
	return e.del
}

func (e *Engine) Fanout() <-chan time.Time {
	return e.fan
}
