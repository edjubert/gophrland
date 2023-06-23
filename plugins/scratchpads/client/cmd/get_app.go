package cmd

import (
	"fmt"
	IPC "github.com/edjubert/hyprland-ipc-go"
	"github.com/spf13/cobra"
)

const UnixSocketName = ".gophrland.sock"

func GetToggle(cmd *cobra.Command, args []string) error {
	conn := IPC.StartUnixConnection(UnixSocketName)

	if len(args) != 1 {
		return cmd.Help()
	}

	if _, err := conn.Write([]byte(fmt.Sprintf("scratchpads toggle %s", args[0]))); err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("[ERROR] - Error reading:", err.Error())
		panic(err)
	}

	return nil
}
