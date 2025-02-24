package router

import (
	"time"
)

type Engine struct {
	uid <-chan Packet
	rac <-chan Packet
	tur <-chan Packet

	tic <-chan time.Time
}

func (e *Engine) Uuid() <-chan Packet {
	return e.uid
}

func (e *Engine) Race() <-chan Packet {
	return e.rac
}

func (e *Engine) Turn() <-chan Packet {
	return e.tur
}

func (e *Engine) Tick() <-chan time.Time {
	return e.tic
}
