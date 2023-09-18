package cmd

import "fmt"

type MonitorsOptions struct {
	Workspaces bool `yaml:"workspaces"`
}

const Name = "monitors"

func startWorkspacesPerMonitor() {
	fmt.Println("Hello workspace per monitor")
}

func LoadPlugin(cfg MonitorsOptions) {
	if cfg.Workspaces {
		startWorkspacesPerMonitor()
	}
}
