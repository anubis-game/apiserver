package cmd

import (
	"github.com/anubis-game/apiserver/cmd/daemon"
	"github.com/anubis-game/apiserver/cmd/release"
	"github.com/anubis-game/apiserver/cmd/version"
	"github.com/anubis-game/apiserver/cmd/wallet"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

var (
	use = "apiserver"
	sho = "Golang based apiserver."
	lon = "Golang based apiserver."
)

func New() (*cobra.Command, error) {
	var err error

	// --------------------------------------------------------------------- //

	var dae *cobra.Command
	{
		c := daemon.Config{}

		dae, err = daemon.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var rel *cobra.Command
	{
		c := release.Config{}

		rel, err = release.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var ver *cobra.Command
	{
		c := version.Config{}

		ver, err = version.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	var wal *cobra.Command
	{
		c := wallet.Config{}

		wal, err = wallet.New(c)
		if err != nil {
			return nil, tracer.Mask(err)
		}
	}

	// --------------------------------------------------------------------- //

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
			// We slience errors because we do not want to see spf13/cobra printing.
			// The errors returned by the commands will be propagated to the main.go
			// anyway, where we have custom error printing for the command line
			// tool.
			SilenceErrors: true,
			SilenceUsage:  true,
		}
	}

	{
		c.SetHelpCommand(&cobra.Command{Hidden: true})
	}

	{
		c.AddCommand(dae)
		c.AddCommand(rel)
		c.AddCommand(ver)
		c.AddCommand(wal)
	}

	return c, nil
}
