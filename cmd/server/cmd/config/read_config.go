package config

import (
	"fmt"
	"gophrland/cmd/server/cmd/plugins/expose"
	"gophrland/cmd/server/cmd/plugins/scratchpads"
	"gopkg.in/yaml.v3"
	"os"
)

type Options struct {
	Scratchpads []map[string]scratchpads.Scratchpad `yaml:"scratchpads"`
	Expose      []map[string]expose.Expose          `yaml:"expose"`
}

type Config struct {
	Plugins []string `yaml:"plugins"`
	Options Options  `yaml:"options"`
}

func ReadConfig(file string) Config {
	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Could not read file '%s' -> %v\n", file, err)
	}

	var config Config
	if err := yaml.Unmarshal(dat, &config); err != nil {
		fmt.Printf("Could not unmarshal %v\n", err)
	}

	return config
}