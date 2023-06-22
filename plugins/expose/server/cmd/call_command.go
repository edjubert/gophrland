package cmd

import "fmt"

const (
	Name      = "expose"
	ToggleCmd = "toggle"
	ShowCmd   = "show"
)

func Command(cmd string, opts ExposeOptions) error {
	fmt.Println("[CMD]: ", cmd)
	switch cmd {
	case ToggleCmd:
		return toggle(opts)
	case ShowCmd:
		return show(opts)
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
