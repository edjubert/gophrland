package cmd

import "fmt"

const (
	Move  = "move"
	Focus = "focus"
)

func Command(cmd string, args []string, opts MonitorsOptions) error {
	fmt.Println("cmd", cmd)
	fmt.Println("args", args)
	switch cmd {
	case Move:
		return move(args, opts)
	case Focus:
		return focus(args, opts)
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
