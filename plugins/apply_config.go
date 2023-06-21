package plugins

import (
	"fmt"
	expose "github.com/edjubert/gophrland/plugins/expose/server/cmd"
	float "github.com/edjubert/gophrland/plugins/float/server/cmd"
	scratchpads "github.com/edjubert/gophrland/plugins/scratchpads/server/cmd"
)

type Options struct {
	Scratchpads []map[string]scratchpads.ScratchpadOptions `yaml:"scratchpads"`
	Expose      expose.ExposeOptions                       `yaml:"expose"`
	Float       float.BringFloatOptions                    `yaml:"float"`
}

type Config struct {
	Plugins []string `yaml:"plugins"`
	Options Options  `yaml:"options"`
}

const (
	Scratchpads = "scratchpads"
	Expose      = "expose"
	Float       = "float"
)

func ApplyConfig(config Config) {
	for _, plugin := range config.Plugins {
		switch plugin {
		case Scratchpads:
			if err := scratchpads.LoadPlugin(config.Options.Scratchpads); err != nil {
				fmt.Printf("[ERROR] - loading plugin %s", plugin)
				return
			}
		case Expose:
			expose.LoadPlugin()
		case Float:
			float.LoadPlugin()
		default:
			fmt.Printf("[WARN] - plugin '%s' is not implemented yet\n", plugin)
		}
	}
}
