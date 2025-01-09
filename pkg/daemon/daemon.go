package daemon

import (
	"net"
	"time"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/anubis-game/apiserver/pkg/stream"
	"github.com/gorilla/mux"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Don chan struct{}
	Env envvar.Env
}

type Daemon struct {
	env envvar.Env
	lis net.Listener
	log logger.Interface
	rtr *mux.Router
	str *stream.Stream
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

	var rtr *mux.Router
	{
		rtr = mux.NewRouter()
	}

	var str *stream.Stream
	{
		str = stream.New(stream.Config{
			Don: c.Don,
			Log: log,
			Out: musDur(c.Env.ConnectionTimeout, "s"),
			Reg: reg,
		})
	}

	return &Daemon{
		env: c.Env,
		lis: lis,
		log: log,
		rtr: rtr,
		str: str,
	}
}

func musDur(str string, uni string) time.Duration {
	dur, err := time.ParseDuration(str + uni)
	if err != nil {
		tracer.Panic(err)
	}

	return dur
}
