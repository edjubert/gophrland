package cmd

import "fmt"

const (
	TOGGLE = "toggle"
	SHOW   = "show"
)

func Command(cmd string, opts ExposeOptions) error {
	fmt.Println("[CMD]: ", cmd)
	switch cmd {
	case TOGGLE:
		return toggle(opts)
	case SHOW:
		return show(opts)
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
