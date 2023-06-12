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
	for _, plugin := range config.Plugins {
		switch plugin {
		case SCRATCHPADS:
			if err := scratchpads.LoadPlugin(config.Options.Scratchpads); err != nil {
				fmt.Printf("[ERROR] - loading plugin %s", plugin)
				return
			}
		case EXPOSE:
			expose.LoadPlugin()
		default:
			fmt.Printf("[WARN] - plugin '%s' is not implemented yet\n", plugin)
		}
	}
}
