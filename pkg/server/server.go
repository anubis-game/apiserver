package server

import (
	"fmt"
	"net"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/matrix"
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
	Don chan struct{}
	Env envvar.Env
	Lis net.Listener
	Log logger.Interface
	Reg *registry.Registry
	Rtr *router.Client
	Tkx *tokenx.TokenX[common.Address]
	Uni *unique.Unique[common.Address, byte]
}

type Server struct {
	don chan struct{}
	env envvar.Env
	lis net.Listener
	log logger.Interface

	agl *random.Random
	con *connect.Handler
	crd *random.Random
	qdr *random.Random
	rtr *mux.Router
}

func New(c Config) *Server {
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Lis == nil {
		tracer.Panic(fmt.Errorf("%T.Lis must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}

	var agl *random.Random
	{
		agl = random.New(random.Config{
			Buf: c.Env.EngineCapacity,
			Don: c.Don,
			Log: c.Log,
			Max: 255,
			Min: 0,
		})
	}

	// We generate random coordinates based on a threshold around the edges of the
	// game map, in which players cannot be placed initially, upon joining the
	// game. The purpose of this buffer region is to not put players too close to
	// the edges of the game, so that they cannot run into the wall accidentally.

	var crd *random.Random
	{
		crd = random.New(random.Config{
			Buf: c.Env.EngineCapacity * 2,
			Don: c.Don,
			Log: c.Log,
			Max: matrix.Max - int(matrix.Pt8),
			Min: matrix.Min + int(matrix.Pt8),
		})
	}

	var qdr *random.Random
	{
		qdr = random.New(random.Config{
			Buf: c.Env.EngineCapacity,
			Don: c.Don,
			Log: c.Log,
			Max: 4,
			Min: 1,
		})
	}

	var con *connect.Handler
	{
		con = connect.New(connect.Config{
			Agl: agl,
			Crd: crd,
			Qdr: qdr,

			Don: c.Don,
			Log: c.Log,
			Reg: c.Reg,
			Rtr: c.Rtr,
			Tkx: c.Tkx,
			Uni: c.Uni,

			Cap: c.Env.EngineCapacity,
		})
	}

	var rtr *mux.Router
	{
		rtr = mux.NewRouter()
	}

	return &Server{
		don: c.Don,
		env: c.Env,
		lis: c.Lis,
		log: c.Log,

		agl: agl,
		con: con,
		crd: crd,
		qdr: qdr,
		rtr: rtr,
	}
}
