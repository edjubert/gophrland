package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/cmd/gophrland/cmd/server"
	client "github.com/edjubert/gophrland/plugins"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:           filepath.Base(os.Args[0]),
		Short:         "A simple Hyprland window manager",
		Long:          "A simple Hyprland window manager in Golang",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	client.AddCommand(rootCmd)
	server.AddCommand(rootCmd, cfgFile)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home + "/.config/hypr")
		viper.SetConfigType("yaml")
		viper.SetConfigName("gophrland")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("[ERROR] - %v", err)
	}
}
