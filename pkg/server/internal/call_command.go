package internal

import (
	"fmt"
	"github.com/edjubert/gophrland/plugins"
	expose "github.com/edjubert/gophrland/plugins/expose/server/cmd"
	float "github.com/edjubert/gophrland/plugins/float/server/cmd"
	scratchpads "github.com/edjubert/gophrland/plugins/scratchpads/server/cmd"
	"strings"
)

func callCommand(command string, opts plugins.Options) error {
	fields := strings.Fields(command)
	plugin := fields[0]

	switch plugin {
	case plugins.Scratchpads:
		cmd := fields[1]
		args := fields[2:]
		return scratchpads.Command(cmd, args, opts.Scratchpads)
	case plugins.Expose:
		cmd := ""
		if len(fields[1:]) > 0 {
			cmd = fields[1]
		}
		return expose.Command(cmd, opts.Expose)
	case plugins.Float:
		cmd := fields[1]
		return float.Command(cmd, opts.Float)
	default:
		return fmt.Errorf("[ERROR] - %s is not a recognized command\n", plugin)
	}
}
