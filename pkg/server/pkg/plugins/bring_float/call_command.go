package bring_float

import "fmt"

const (
	CURRENT = "current"
)

func Command(cmd string, opts BringFloatOptions) error {
	switch cmd {
	case CURRENT:
		return bringCurrent(opts)
	default:
		return fmt.Errorf("[WARN] - unrecognized command")
	}
}
