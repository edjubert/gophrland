package cmd

import (
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
)

const EXPOSED_SPECIAL_WORKSPACE = "exposed_special_workspace"

func show(options ExposeOptions) error {
	name := EXPOSED_SPECIAL_WORKSPACE
	if options.Name != "" {
		name = options.Name
	}
	d := hyprctl.Dispatch{}
	return d.Toggle.SpecialWorkspace(name)
}
