package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/pkg/client/pkg/tools"
	"github.com/spf13/cobra"
)

func GetBringCurrent(_ *cobra.Command, args []string) error {
	conn := tools.StartTCPConnection("localhost", 9988)

	if len(args) > 0 {
		return fmt.Errorf("[ERROR] - this command do not take arguments")
	}

	if _, err := conn.Write([]byte("float bring current")); err != nil {
		return fmt.Errorf("[ERROR] - %w\n", err)
	}

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("[ERROR] - %w\n", err)
	}

	return nil
}
