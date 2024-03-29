package internal

import (
	"fmt"
	"github.com/edjubert/gophrland/plugins"
	expose "github.com/edjubert/gophrland/plugins/expose/server/cmd"
	float "github.com/edjubert/gophrland/plugins/float/server/cmd"
	monitors "github.com/edjubert/gophrland/plugins/monitors/server/cmd"
	scratchpads "github.com/edjubert/gophrland/plugins/scratchpads/server/cmd"
	"strings"
)

func callCommand(command string, opts plugins.Options) error {
	fmt.Println("command: ", command)
	fields := strings.Fields(command)
	plugin := fields[0]

	switch plugin {
	case scratchpads.Name:
		cmd := fields[1]
		args := fields[2:]
		return scratchpads.Command(cmd, args, opts.Scratchpads)
	case expose.Name:
		cmd := ""
		if len(fields[1:]) > 0 {
			cmd = fields[1]
		}
		return expose.Command(cmd, opts.Expose)
	case float.Name:
		cmd := fields[1]
		return float.Command(cmd, opts.Float)
	case monitors.Name:
		cmd := fields[1]
		args := fields[2:]
		return monitors.Command(cmd, args, opts.Monitors)
	default:
		return fmt.Errorf("[ERROR] - %s is not a recognized command\n", plugin)
	}
}
