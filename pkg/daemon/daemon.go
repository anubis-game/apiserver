package daemon

import (
	"net"

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
	don chan struct{}
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

	var rtr *mux.Router
	{
		rtr = mux.NewRouter()
	}

	var str *stream.Stream
	{
		str = stream.New(stream.Config{
			Log: log,
		})
	}

	return &Daemon{
		don: c.Don,
		env: c.Env,
		lis: lis,
		log: log,
		rtr: rtr,
		str: str,
	}
}
