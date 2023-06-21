package cmd

import "fmt"

const (
	Bring = "bring"
)

func Command(cmd string, opts BringFloatOptions) error {
	switch cmd {
	case Bring:
		return bringCurrent(opts)
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
