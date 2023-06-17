package expose

import "fmt"

const (
	TOGGLE = "toggle"
	ROOT   = ""
)

func Command(cmd string, opts ExposeOptions) error {
	switch cmd {
	case TOGGLE:
		return toggle(opts)
	case ROOT:
		return root(opts)
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
