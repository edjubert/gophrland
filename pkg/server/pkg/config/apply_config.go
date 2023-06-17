package config

import (
	"fmt"
	"gophrland/pkg/server/pkg/plugins/bring_float"
	"gophrland/pkg/server/pkg/plugins/expose"
	"gophrland/pkg/server/pkg/plugins/scratchpads"
)

const (
	SCRATCHPADS = "scratchpads"
	EXPOSE      = "expose"
	BRING_FLOAT = "bring_float"
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
		case BRING_FLOAT:
			bring_float.LoadPlugin()
		default:
			fmt.Printf("[WARN] - plugin '%s' is not implemented yet\n", plugin)
		}
	}
}
