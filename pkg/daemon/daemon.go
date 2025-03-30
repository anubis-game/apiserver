package daemon

import (
	"fmt"
	"math"
	"net"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/engine"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/anubis-game/apiserver/pkg/random"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/server"
	"github.com/anubis-game/apiserver/pkg/tokenx"
	"github.com/anubis-game/apiserver/pkg/unique"
	"github.com/anubis-game/apiserver/pkg/worker"
	"github.com/anubis-game/apiserver/pkg/worker/release"
	"github.com/anubis-game/apiserver/pkg/worker/resolve"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don chan struct{}
	Env envvar.Env
}

type Daemon struct {
	eng *engine.Engine
	lis net.Listener
	log logger.Interface
	reg *registry.Registry
	ser *server.Server
	tkx *tokenx.TokenX[common.Address]
	wrk *worker.Worker
}

func New(c Config) *Daemon {
	var err error

	if c.Env.EngineCapacity > math.MaxUint8 {
		tracer.Panic(fmt.Errorf("%T.Env.EngineCapacity must not be larger than max byte", c))
	}

	var log logger.Interface
	{
		log = logger.New(logger.Config{
			Filter: logger.NewLevelFilter(c.Env.LogLevel),
		})
	}

	var lis net.Listener
	{
		lis, err = net.Listen("tcp", net.JoinHostPort(c.Env.HttpHost, c.Env.HttpPort))
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}

	var agl *random.Random
	{
		agl = random.New(random.Config{
			Buf: c.Env.EngineCapacity,
			Don: c.Don,
			Log: log,
			Max: 255,
			Min: 0,
		})
	}

	// We generate random coordinates based on a threshold around the edges of the
	// game map, in which players cannot be placed initially, upon joining the
	// game. There are two benefits to this restriction.
	//
	//     1. Players cannot run into the wall accidentally upon joining the game.
	//
	//     2. We prevent negative coordinate values during internal player lookups.
	//

	var crd *random.Random
	{
		crd = random.New(random.Config{
			Buf: c.Env.EngineCapacity * 4,
			Don: c.Don,
			Log: log,
			Max: matrix.Max - int(2*matrix.Pt8),
			Min: matrix.Min + int(2*matrix.Pt8),
		})
	}

	var qdr *random.Random
	{
		qdr = random.New(random.Config{
			Buf: c.Env.EngineCapacity,
			Don: c.Don,
			Log: log,
			Max: 4,
			Min: 1,
		})
	}

	var reg *registry.Registry
	{
		reg = registry.New(registry.Config{
			Add: c.Env.ChainRegistryContract,
			Key: c.Env.SignerPrivateKey,
			Log: log,
			RPC: c.Env.ChainRpcEndpoint,
		})
	}

	var wrk *worker.Worker
	{
		wrk = worker.New(worker.Config{
			Don: c.Don,
			Log: log,
			Reg: reg,
			Sig: []worker.Signer{
				release.New(release.Config{Reg: reg}),
				resolve.New(resolve.Config{Reg: reg}),
			},
		})
	}

	var rtr *router.Router
	{
		rtr = router.New(router.Config{
			Cap: c.Env.EngineCapacity,
		})
	}

	var tkx *tokenx.TokenX[common.Address]
	{
		tkx = tokenx.New[common.Address]()
	}

	var uni *unique.Unique[common.Address, byte]
	{
		uni = unique.New[common.Address, byte]()
	}

	var eng *engine.Engine
	{
		eng = engine.New(engine.Config{
			Crd: crd,
			Cap: c.Env.EngineCapacity,
			Don: c.Don,
			Log: log,
			Rtr: rtr.Engine(),
			Tkx: tkx,
			Uni: uni,
			Wrk: wrk,
		})
	}

	var ser *server.Server
	{
		ser = server.New(server.Config{
			Agl: agl,
			Crd: crd,
			Don: c.Don,
			Env: c.Env,
			Lis: lis,
			Log: log,
			Qdr: qdr,
			Reg: reg,
			Rtr: rtr.Client(),
			Tkx: tkx,
			Uni: uni,
		})
	}

	return &Daemon{
		eng: eng,
		lis: lis,
		log: log,
		reg: reg,
		ser: ser,
		tkx: tkx,
		wrk: wrk,
	}
}
