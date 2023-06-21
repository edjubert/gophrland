package cmd

import (
	"github.com/spf13/cobra"
)

func getBringCurrent() *cobra.Command {
	return &cobra.Command{
		Use:   "current",
		Short: "Bring all floating windows to center if out of window for current workspace",
		Long:  "Bring all floating windows to center if out of window for current workspace",
		RunE:  GetBringCurrent,
	}
}

func getBring() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bring",
		Short: "Bring all floating windows to center if out of window",
		Long:  "Bring all floating windows to center if out of window",
	}

	cmd.AddCommand(getBringCurrent())
	return cmd
}

func Float() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "float",
		Short: "Bring Float CLI handler",
		Long:  "These are the tools to use the 'bring_float' plugin with the CLI",
	}

	cmd.AddCommand(getBring())
	return cmd
}
