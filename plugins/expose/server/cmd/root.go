package cmd

import (
	IPC "github.com/edjubert/hyprland-ipc-go"
)

const EXPOSED_SPECIAL_WORKSPACE = "exposed_special_workspace"

func show(options ExposeOptions) error {
	name := EXPOSED_SPECIAL_WORKSPACE
	if options.Name != "" {
		name = options.Name
	}
	return IPC.ToggleSpecialWorkspace(name)
}
