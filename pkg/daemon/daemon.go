package daemon

import (
	"fmt"
	"math"
	"net"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/engine"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/filler"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/server"
	"github.com/anubis-game/apiserver/pkg/server/handler/connect"
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
	con *connect.Handler
	eng *engine.Engine
	fil *filler.Filler
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
		tracer.Panic(fmt.Errorf("c.Env.EngineCapacity must not be larger than max byte"))
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

	var con *connect.Handler
	{
		con = connect.New(connect.Config{
			Cap: c.Env.EngineCapacity,
			Don: c.Don,
			Log: log,
			Reg: reg,
			Rtr: rtr.Client(),
			Tkx: tkx,
			Uni: uni,
		})
	}

	var fil *filler.Filler
	{
		fil = filler.New(filler.Config{
			Cap: c.Env.EngineCapacity,
			Don: c.Don,
			Log: log,
		})
	}

	var eng *engine.Engine
	{
		eng = engine.New(engine.Config{
			Cap: c.Env.EngineCapacity,
			Don: c.Don,
			Fil: fil,
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
			Con: con,
			Don: c.Don,
			Env: c.Env,
			Lis: lis,
			Log: log,
		})
	}

	return &Daemon{
		con: con,
		eng: eng,
		fil: fil,
		lis: lis,
		log: log,
		reg: reg,
		ser: ser,
		tkx: tkx,
		wrk: wrk,
	}
}
