package stream

import (
	"time"

	"github.com/anubis-game/apiserver/pkg/client"
)

type Packet struct {
	Byt []byte
	Cli *client.Client
}

type Router struct {
	Closer <-chan struct{}
	Create chan Packet
	Delete chan Packet
	Update <-chan time.Time
	// Auth Action
	// Move Action
	// Kill Action
}

func NewRouter(don <-chan struct{}) *Router {
	return &Router{
		Closer: don,
		Create: make(chan Packet, 500),
		Delete: make(chan Packet, 500),
		Update: time.NewTicker(25 * time.Millisecond).C,
	}
}
