package cmd

import (
	"github.com/spf13/cobra"
)

func getExposeToggleCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "toggle",
		Short: "Toggle visibility of expose",
		Long:  "Toggle visibility of expose",
		RunE:  GetToggle,
	}
}

func getExposeShowCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Toggle visibility of expose special workspace",
		Long:  "Toggle visibility of expose special workspace",
		RunE:  GetShow,
	}
}

func Expose() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expose",
		Short: "Expose CLI handler",
		Long:  "These are tools to use the expose plugin from the CLI",
	}

	cmd.AddCommand(getExposeToggleCommand())
	cmd.AddCommand(getExposeShowCommand())

	return cmd
}
