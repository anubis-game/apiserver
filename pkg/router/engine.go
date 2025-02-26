package router

import (
	"time"
)

type Engine struct {
	rac <-chan byte
	tur <-chan Turn
	uid <-chan Uuid

	tic <-chan time.Time
}

func (e *Engine) Race() <-chan byte {
	return e.rac
}

func (e *Engine) Turn() <-chan Turn {
	return e.tur
}

func (e *Engine) Uuid() <-chan Uuid {
	return e.uid
}

func (e *Engine) Tick() <-chan time.Time {
	return e.tic
}
