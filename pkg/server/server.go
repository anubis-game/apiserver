package server

import (
	"fmt"
	"net"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/random"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/server/handler/connect"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gorilla/mux"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Agl *random.Random
	Crd *random.Random
	Don chan struct{}
	Env envvar.Env
	Lis net.Listener
	Log logger.Interface
	Qdr *random.Random
	Reg *registry.Registry
	Rtr *router.Client
	Tkx *tokenx.TokenX[common.Address]
	Uni *unique.Unique[common.Address, byte]
}

type Server struct {
	agl *random.Random
	con *connect.Handler
	crd *random.Random
	don chan struct{}
	env envvar.Env
	lis net.Listener
	log logger.Interface
	qdr *random.Random
	rtr *mux.Router
}

func New(c Config) *Server {
	if c.Agl == nil {
		tracer.Panic(fmt.Errorf("%T.Agl must not be empty", c))
	}
	if c.Crd == nil {
		tracer.Panic(fmt.Errorf("%T.Crd must not be empty", c))
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
	if c.Qdr == nil {
		tracer.Panic(fmt.Errorf("%T.Qdr must not be empty", c))
	}

	var con *connect.Handler
	{
		con = connect.New(connect.Config{
			Agl: c.Agl,
			Cap: c.Env.EngineCapacity,
			Crd: c.Crd,
			Don: c.Don,
			Log: c.Log,
			Qdr: c.Qdr,
			Reg: c.Reg,
			Rtr: c.Rtr,
			Tkx: c.Tkx,
			Uni: c.Uni,
		})
	}

	var rtr *mux.Router
	{
		rtr = mux.NewRouter()
	}

	return &Server{
		agl: c.Agl,
		con: con,
		crd: c.Crd,
		don: c.Don,
		env: c.Env,
		lis: c.Lis,
		log: c.Log,
		qdr: c.Qdr,
		rtr: rtr,
	}
}
