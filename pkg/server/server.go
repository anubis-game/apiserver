package server

import (
	"fmt"
	"net"

	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/stream"
	"github.com/gorilla/mux"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Env envvar.Env
	Lis net.Listener
	Log logger.Interface
	Rtr *mux.Router
	Str *stream.Stream
}

type Server struct {
	env envvar.Env
	lis net.Listener
	log logger.Interface
	rtr *mux.Router
	str *stream.Stream
}

func New(c Config) *Server {
	if c.Lis == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Lis must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.Rtr == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Rtr must not be empty", c)))
	}
	if c.Str == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Str must not be empty", c)))
	}

	return &Server{
		env: c.Env,
		lis: c.Lis,
		log: c.Log,
		rtr: c.Rtr,
		str: c.Str,
	}
}
