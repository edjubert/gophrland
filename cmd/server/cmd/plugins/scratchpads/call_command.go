package scratchpads

import (
	"fmt"
)

const (
	TOGGLE = "toggle"
	SHOW   = "show"
	HIDE   = "hide"
)

func Command(cmd string, args []string) error {
	switch cmd {
	case TOGGLE:
		return toggle(args)
	case SHOW, HIDE:
		return fmt.Errorf("not (yet) implemented")
	default:
		return fmt.Errorf("unrecognized command")
	}
}
