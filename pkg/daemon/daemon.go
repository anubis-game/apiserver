package daemon

import (
	"net"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/engine"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/router"
	"github.com/anubis-game/apiserver/pkg/server"
	"github.com/anubis-game/apiserver/pkg/server/handler/connect"
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
	lis net.Listener
	log logger.Interface
	reg *registry.Registry
	rel *worker.Worker[common.Address, release.Packet]
	res *worker.Worker[common.Address, resolve.Packet]
	ser *server.Server
}

func New(c Config) *Daemon {
	var err error

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

	var rel *worker.Worker[common.Address, release.Packet]
	{
		rel = newRel(c.Don, log, reg)
	}

	var res *worker.Worker[common.Address, resolve.Packet]
	{
		res = newRes(c.Don, log, reg)
	}

	var rtr *router.Router
	{
		rtr = router.New()
	}

	var con *connect.Handler
	{
		con = connect.New(connect.Config{
			Don: c.Don,
			Env: c.Env,
			Log: log,
			Reg: reg,
			Rel: rel,
			Res: res,
			Rtr: rtr.Client(),
		})
	}

	var eng *engine.Engine
	{
		eng = engine.New(engine.Config{
			Don: c.Don,
			Log: log,
			Rtr: rtr.Engine(),
		})
	}

	var ser *server.Server
	{
		ser = server.New(server.Config{
			Con: con,
			Env: c.Env,
			Lis: lis,
			Log: log,
		})
	}

	return &Daemon{
		con: con,
		eng: eng,
		lis: lis,
		log: log,
		reg: reg,
		rel: rel,
		res: res,
		ser: ser,
	}
}

func newRel(don <-chan struct{}, log logger.Interface, reg *registry.Registry) *worker.Worker[common.Address, release.Packet] {
	var rel *release.Release
	{
		rel = release.New(release.Config{
			Log: log,
			Reg: reg,
		})
	}

	var wrk *worker.Worker[common.Address, release.Packet]
	{
		wrk = worker.New(worker.Config[common.Address, release.Packet]{
			Don: don,
			Ens: rel,
		})
	}

	return wrk
}

func newRes(don <-chan struct{}, log logger.Interface, reg *registry.Registry) *worker.Worker[common.Address, resolve.Packet] {
	var res *resolve.Resolve
	{
		res = resolve.New(resolve.Config{
			Log: log,
			Reg: reg,
		})
	}

	var wrk *worker.Worker[common.Address, resolve.Packet]
	{
		wrk = worker.New(worker.Config[common.Address, resolve.Packet]{
			Don: don,
			Ens: res,
		})
	}

	return wrk
}
