package version

import (
	"fmt"
	"os"

	"github.com/anubis-game/apiserver/pkg/runtime"
	"github.com/spf13/cobra"
)

type run struct{}

func (r *run) run(cmd *cobra.Command, arg []string) {
	fmt.Fprintf(os.Stdout, "%s\n", runtime.Json())
}
