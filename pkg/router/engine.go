package router

import (
	"time"
)

type Engine struct {
	joi <-chan Packet
	mov <-chan Packet
	rac <-chan Packet

	psh <-chan time.Time
}

func (e *Engine) Join() <-chan Packet {
	return e.joi
}

func (e *Engine) Move() <-chan Packet {
	return e.mov
}

func (e *Engine) Race() <-chan Packet {
	return e.rac
}

func (e *Engine) Push() <-chan time.Time {
	return e.psh
}
