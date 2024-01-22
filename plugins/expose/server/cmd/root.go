package cmd

import (
	"github.com/edjubert/hyprland-ipc-go/hyprctl/dispatch"
)

const EXPOSED_SPECIAL_WORKSPACE = "exposed_special_workspace"

func show(options ExposeOptions) error {
	name := EXPOSED_SPECIAL_WORKSPACE
	if options.Name != "" {
		name = options.Name
	}
	toggle := dispatch.Toggle{}
	return toggle.SpecialWorkspace(name)
}
