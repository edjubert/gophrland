package plugins

import (
	"fmt"

	expose "github.com/edjubert/gophrland/plugins/expose/server/cmd"
	float "github.com/edjubert/gophrland/plugins/float/server/cmd"
	parallax "github.com/edjubert/gophrland/plugins/parallax/server/cmd"
	scratchpads "github.com/edjubert/gophrland/plugins/scratchpads/server/cmd"
)

type Options struct {
	Scratchpads []map[string]scratchpads.ScratchpadOptions `yaml:"scratchpads"`
	Expose      expose.ExposeOptions                       `yaml:"expose"`
	Float       float.BringFloatOptions                    `yaml:"float"`
	Parallax    parallax.ParallaxOptions                   `yaml:"parallax"`
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
				fmt.Printf("[ERROR] - loading plugin %s", plugin)
				return
			}
		case expose.Name:
			expose.LoadPlugin()
		case float.Name:
			float.LoadPlugin()
		case parallax.Name:
			parallax.LoadPlugin()
		default:
			fmt.Printf("[WARN] - plugin '%s' is not implemented yet\n", plugin)
		}
	}
}
