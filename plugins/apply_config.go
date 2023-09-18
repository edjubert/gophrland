package plugins

import (
	"fmt"
	expose "github.com/edjubert/gophrland/plugins/expose/server/cmd"
	float "github.com/edjubert/gophrland/plugins/float/server/cmd"
	monitors "github.com/edjubert/gophrland/plugins/monitors/server/cmd"
	scratchpads "github.com/edjubert/gophrland/plugins/scratchpads/server/cmd"
)

type Options struct {
	Scratchpads []map[string]scratchpads.ScratchpadOptions `yaml:"scratchpads"`
	Expose      expose.ExposeOptions                       `yaml:"expose"`
	Float       float.FloatOptions                         `yaml:"float"`
	Monitors    monitors.MonitorsOptions                   `yaml:"monitors"`
}

type Config struct {
	Plugins []string `yaml:"plugins"`
	Options Options  `yaml:"options"`
}

func ApplyConfig(config Config) {
	for _, plugin := range config.Plugins {
		switch plugin {
		case scratchpads.Name:
			if err := scratchpads.LoadPlugin(config.Options.Scratchpads); err != nil {
				fmt.Printf("[ERROR] - loading plugin '%s' -> %v\n", plugin, err)
				return
			}
			go scratchpads.LoadEventCallbacks(config.Options.Scratchpads)
		case expose.Name:
			expose.LoadPlugin()
		case float.Name:
			float.LoadPlugin()
		case monitors.Name:
			monitors.LoadPlugin(config.Options.Monitors)
		default:
			fmt.Printf("[WARN] - plugin '%s' is not implemented yet\n", plugin)
		}
	}
}
