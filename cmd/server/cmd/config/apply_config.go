package config

import (
	"fmt"
	"gophrland/cmd/server/cmd/plugins/expose"
	"gophrland/cmd/server/cmd/plugins/scratchpads"
)

const (
	SCRATCHPADS = "scratchpads"
	EXPOSE      = "expose"
)

func ApplyConfig(config Config) {
	msgTemplate := "[INFO] - Loading plugin '%s'\n"
	for _, plugin := range config.Plugins {
		switch plugin {
		case SCRATCHPADS:
			fmt.Printf(msgTemplate, plugin)
			scratchpads.LoadPlugin(config.Options.Scratchpads)
		case EXPOSE:
			fmt.Printf(msgTemplate, plugin)
			expose.LoadPlugin(config.Options.Expose)
		default:
			fmt.Printf("[WARN] - plugin '%s' is not implemented yet\n", plugin)
		}
	}
}
