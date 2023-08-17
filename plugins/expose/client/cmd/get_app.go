package cmd

import (
	"fmt"
	IPC "github.com/edjubert/hyprland-ipc-go/ipc"
	"github.com/spf13/cobra"
)

const UnixSocketName = ".gophrland.sock"

func GetShow(_ *cobra.Command, args []string) error {
	conn := IPC.StartUnixConnection(UnixSocketName)

	if len(args) != 0 {
		return fmt.Errorf("[ERROR] -  no arg necessary")
	}

	if _, err := conn.Write([]byte("expose show")); err != nil {
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

func GetToggle(_ *cobra.Command, args []string) error {
	conn := IPC.StartUnixConnection(UnixSocketName)
	if len(args) != 0 {
		return fmt.Errorf("[ERROR] -  no arg necessary")
	}

	if _, err := conn.Write([]byte("expose toggle")); err != nil {
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
