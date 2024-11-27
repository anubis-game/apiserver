package wallet

import (
	"github.com/spf13/cobra"
)

const (
	use = "wallet"
	sho = "Generate Ethereum wallets."
	lon = "Generate Ethereum wallets."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			RunE:  (&run{}).runE,
		}
	}

	return c, nil
}
