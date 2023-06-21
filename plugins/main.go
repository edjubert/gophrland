package plugins

import (
	expose "github.com/edjubert/gophrland/plugins/expose/client/cmd"
	float "github.com/edjubert/gophrland/plugins/float/client/cmd"
	scratchpads "github.com/edjubert/gophrland/plugins/scratchpads/client/cmd"
	"github.com/spf13/cobra"
)

const ServerHost = "localhost"

func AddCommand(cmd *cobra.Command) {
	cmd.AddCommand(scratchpads.Scratchpads())
	cmd.AddCommand(expose.Expose())
	cmd.AddCommand(float.Float())
}