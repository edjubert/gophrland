package cmd

import "fmt"

const (
	Name     = "float"
	BringCmd = "bring"
)

func Command(cmd string, opts BringFloatOptions) error {
	switch cmd {
	case BringCmd:
		return bringCurrent(opts)
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
