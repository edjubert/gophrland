package scratchpads

import (
	"fmt"
)

const (
	TOGGLE = "toggle"
	SHOW   = "show"
	HIDE   = "hide"
)

func Command(cmd string, args []string, opts []map[string]ScratchpadOptions) error {
	switch cmd {
	case TOGGLE:
		return toggle(args, opts)
	case SHOW, HIDE:
		return fmt.Errorf("[ERROR] - not (yet) implemented")
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
