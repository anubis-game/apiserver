package cmd

import (
	"github.com/spf13/cobra"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, arg []string) {
	err := cmd.Help()
	if err != nil {
		panic(err)
	}
}
