package daemon

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anubis-game/apiserver/pkg/daemon"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/spf13/cobra"
)

type run struct {
	flag *flag
}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var env envvar.Env
	{
		env = envvar.Load(r.flag.Env)
	}

	var sig chan os.Signal
	{
		sig = make(chan os.Signal, 2)
	}

	// --------------------------------------------------------------------- //

	var dae *daemon.Daemon
	{
		dae = daemon.New(daemon.Config{
			Env: env,
			Sig: sig,
		})
	}

	{
		go dae.Server().Daemon()
	}

	// --------------------------------------------------------------------- //

	{
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	}

	{
		<-sig
	}

	select {
	case <-time.After(10 * time.Second):
		// One SIGTERM gives the daemon some time to tear down gracefully.
	case <-sig:
		// Two SIGTERMs stop the immediatelly.
	}

	{
		close(sig)
	}

	return nil
}
