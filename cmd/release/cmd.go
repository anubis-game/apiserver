package release

import (
	"github.com/spf13/cobra"
)

const (
	use = "release"
	sho = "Release player addresses for testing."
	lon = "Release player addresses for testing."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var f *flag
	{
		f = &flag{}
	}

	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			RunE:  (&run{flag: f}).runE,
		}
	}

	{
		f.Init(c)
	}

	return c, nil
}
