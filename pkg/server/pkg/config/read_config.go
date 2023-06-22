package config

import (
	"fmt"
	"github.com/edjubert/gophrland/plugins"
	IPC "github.com/edjubert/hyprland-ipc-go"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig(file string) plugins.Config {
	dat, err := os.ReadFile(file)
	if err != nil {
		_ = IPC.SendNotification(5000, "warning", fmt.Sprintf("Could not read config file '%s' -> %v\n", file, err))
		fmt.Printf("[ERROR] - Could not read file '%s' -> %v\n", file, err)
	}

	var config plugins.Config
	if err := yaml.Unmarshal(dat, &config); err != nil {
		fmt.Printf("[ERROR] - Could not unmarshal %v\n", err)
	}

	return config
}
