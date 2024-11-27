package daemon

import (
	"fmt"

	"github.com/spf13/cobra"
)

type flag struct {
	Env string
}

func (f *flag) Init(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Env, "env", "local", "the env file to load, e.g. sepolia for env.sepolia")
}

func (f *flag) Validate() error {
	if f.Env != "local" && f.Env != "sepolia" {
		return fmt.Errorf("--env must be either local or sepolia")
	}

	return nil
}
