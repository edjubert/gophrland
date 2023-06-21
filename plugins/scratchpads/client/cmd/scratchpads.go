package cmd

import (
	"github.com/spf13/cobra"
)

func getScratchpadsToggleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "toggle",
		Short: "Toggle visibility of your scratchpads",
		Long:  "Toggle visibility of your scratchpads by running animation",
		RunE:  GetToggle,
	}
}

func Scratchpads() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scratchpads",
		Short: "Tools to handle scratchpads",
		Long: `To use, activate the plugin in your 'gophrland.yaml' and add scratchpads:
plugins:
  - "scratchpads"

options:
  scratchpads:
    - myScratchpad:
      command: "my command"
      animation: "fromTop"
      unfocus: "hide"
      margin: 60`,
		ValidArgs: []string{"toggle"},
	}

	cmd.AddCommand(getScratchpadsToggleCommand())
	return cmd
}
