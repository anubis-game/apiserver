package router

import (
	"time"
)

type Engine struct {
	uid <-chan Packet
	mov <-chan Packet
	rac <-chan Packet

	tic <-chan time.Time
}

func (e *Engine) Uuid() <-chan Packet {
	return e.uid
}

func (e *Engine) Move() <-chan Packet {
	return e.mov
}

func (e *Engine) Race() <-chan Packet {
	return e.rac
}

func (e *Engine) Tick() <-chan time.Time {
	return e.tic
}
