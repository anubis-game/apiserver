package server

import (
	"fmt"
	"net"

	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/server/handler/connect"
	"github.com/gorilla/mux"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Con *connect.Handler
	Don chan struct{}
	Env envvar.Env
	Lis net.Listener
	Log logger.Interface
}

type Server struct {
	con *connect.Handler
	don chan struct{}
	env envvar.Env
	lis net.Listener
	log logger.Interface
	rtr *mux.Router
}

func New(c Config) *Server {
	if c.Con == nil {
		tracer.Panic(fmt.Errorf("%T.Con must not be empty", c))
	}
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Lis == nil {
		tracer.Panic(fmt.Errorf("%T.Lis must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}

	return &Server{
		con: c.Con,
		don: c.Don,
		env: c.Env,
		lis: c.Lis,
		log: c.Log,
		rtr: mux.NewRouter(),
	}
}
