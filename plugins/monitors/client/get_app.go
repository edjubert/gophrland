package client

import (
	"fmt"
	monitors "github.com/edjubert/gophrland/plugins/monitors/server/cmd"
	IPC "github.com/edjubert/hyprland-ipc-go/ipc"
	"github.com/spf13/cobra"
)

const (
	FocusCmd       = "focus"
	MoveCmd        = "move"
	NextArg        = "next"
	PrevArg        = "previous"
	UnixSocketName = ".gophrland.sock"
)

func GetCommand(ccmd string, args string) func(cmd *cobra.Command, _ []string) error {
	return func(cmd *cobra.Command, _ []string) error {
		conn := IPC.StartUnixConnection(UnixSocketName)

		acceptedCmds := []string{FocusCmd, MoveCmd}
		if args != NextArg && args != PrevArg {
			return fmt.Errorf("[ERROR] - could not recognize arg -> %s\n")
		}
		switch ccmd {
		case FocusCmd, MoveCmd:
			if _, err := conn.Write([]byte(fmt.Sprintf("%s %s %s", monitors.Name, ccmd, args))); err != nil {
				return err
			}
		default:
			return fmt.Errorf("[ERROR] - could not recognize command, only %v accepted\n", acceptedCmds)
		}

		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("[ERROR] - Error reading: ", err.Error())
			panic(err)
		}

		return err

	}
}
