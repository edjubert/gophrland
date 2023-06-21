package cmd

import (
	"fmt"
	client "github.com/edjubert/gophrland/pkg/client/pkg"
	server "github.com/edjubert/gophrland/pkg/server/pkg"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gophrland",
	Long:  "All software has version. This is Gophrland's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Gophrland Hyprland Window Manager -- server: %s -- client: %s\n", server.Version, client.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
