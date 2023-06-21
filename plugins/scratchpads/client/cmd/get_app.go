package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/pkg/client/pkg/tools"
	"github.com/spf13/cobra"
)

func GetToggle(cmd *cobra.Command, args []string) error {
	conn := tools.StartTCPConnection("localhost", 9988)

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
