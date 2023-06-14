package config

import (
	"fmt"
	"gophrland/cmd/server/cmd/plugins/bring_float"
	"gophrland/cmd/server/cmd/plugins/expose"
	"gophrland/cmd/server/cmd/plugins/scratchpads"
	"gopkg.in/yaml.v3"
	"os"
)

type Options struct {
	Scratchpads []map[string]scratchpads.ScratchpadOptions `yaml:"scratchpads"`
	Expose      expose.ExposeOptions                       `yaml:"expose"`
	BringFloat  bring_float.BringFloatOptions              `yaml:"bring_float"`
}

type Config struct {
	Plugins []string `yaml:"plugins"`
	Options Options  `yaml:"options"`
}

func ReadConfig(file string) Config {
	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("[ERROR] - Could not read file '%s' -> %v\n", file, err)
	}

	var config Config
	if err := yaml.Unmarshal(dat, &config); err != nil {
		fmt.Printf("[ERROR] - Could not unmarshal %v\n", err)
	}

	return config
}
