package cmd

import (
	"fmt"
)

const (
	Name      = "scratchpads"
	ToggleCmd = "toggle"
	ShowCmd   = "show"
	HideCmd   = "hide"
)

func Command(cmd string, args []string, opts []map[string]ScratchpadOptions) error {
	switch cmd {
	case ToggleCmd:
		return toggle(args, opts)
	case ShowCmd, HideCmd:
		return fmt.Errorf("[ERROR] - not (yet) implemented")
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
