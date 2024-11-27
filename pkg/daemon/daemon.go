package daemon

import (
	"os"

	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/xh3b4sd/logger"
)

type Config struct {
	Env envvar.Env
	Sig chan os.Signal
}

type Daemon struct {
	env envvar.Env
	log logger.Interface
	sig chan os.Signal
}

func New(c Config) *Daemon {
	var log logger.Interface
	{
		log = logger.New(logger.Config{
			Filter: logger.NewLevelFilter(c.Env.LogLevel),
		})
	}

	return &Daemon{
		env: c.Env,
		log: log,
		sig: c.Sig,
	}
}
