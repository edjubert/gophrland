package config

import (
	"fmt"
	"github.com/edjubert/gophrland/plugins"
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig(file string) plugins.Config {
	dat, err := os.ReadFile(file)
	if err != nil {
		_ = hyprctl.SendNotification(5000, "warning", fmt.Sprintf("Could not read config file '%s' -> %v\n", file, err))
		fmt.Printf("[ERROR] - Could not read file '%s' -> %v\n", file, err)
	}

	var config plugins.Config
	if err := yaml.Unmarshal(dat, &config); err != nil {
		fmt.Printf("[ERROR] - Could not unmarshal %v\n", err)
	}

	return config
}
