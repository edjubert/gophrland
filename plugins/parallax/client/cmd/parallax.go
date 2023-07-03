package cmd

import (
	"fmt"

	IPC "github.com/edjubert/hyprland-ipc-go"
	"github.com/spf13/cobra"
)

func MyAwesomeFunction(cmd *cobra.Command, args []string) error {
	conn := IPC.StartUnixConnection("")

	// If you command take only one argument
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

func Run() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "parallax",
		Short: "Parallax CLI handler",
		Long:  "These are tools to use the 'parallax' plugin from the CLI",
		RunE:  MyAwesomeFunction,
	}

	return cmd
}
