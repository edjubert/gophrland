package client

import (
	"fmt"
	"github.com/spf13/cobra"
)

func getMonitorsFocusNextCommand(cmd string) *cobra.Command {
	return &cobra.Command{
		Use:     NextArg,
		Aliases: []string{"n"},
		Short:   fmt.Sprintf("%s the next monitor in list (alias: 'n')", cmd),
		Long:    fmt.Sprintf("%s the next monitor in list (alias: 'n')", cmd),
		RunE:    GetCommand(cmd, NextArg),
	}
}

func getMonitorsFocusPrevCommand(cmd string) *cobra.Command {
	return &cobra.Command{
		Use:     PrevArg,
		Aliases: []string{"prev", "p"},
		Short:   "Focus the next monitor in list (alias: 'p', 'prev')",
		Long:    "Focus the next monitor in list (alias: 'p', 'prev')",
		RunE:    GetCommand(cmd, PrevArg),
	}
}

func getMonitorsFocusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       FocusCmd,
		Short:     "Focus the next or previous monitor in monitors list",
		Long:      "Focus the next monitor in monitors list",
		ValidArgs: []string{NextArg, PrevArg},
	}

	cmd.AddCommand(getMonitorsFocusNextCommand(FocusCmd))
	cmd.AddCommand(getMonitorsFocusPrevCommand(FocusCmd))

	return cmd
}
func getMonitorsMoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:       MoveCmd,
		Short:     "Move the next monitor in monitors list",
		Long:      "Move the next monitor in monitors list",
		ValidArgs: []string{NextArg, PrevArg},
	}

	cmd.AddCommand(getMonitorsFocusNextCommand(MoveCmd))
	cmd.AddCommand(getMonitorsFocusPrevCommand(MoveCmd))

	return cmd
}

func Monitors() *cobra.Command {
	cmd := &cobra.Command{
		Use:       "monitors",
		Short:     "Tools to handle monitors movements",
		Long:      "Tools to handle monitors movements",
		ValidArgs: []string{FocusCmd, MoveCmd},
	}

	cmd.AddCommand(getMonitorsFocusCommand())
	cmd.AddCommand(getMonitorsMoveCommand())

	return cmd
}
