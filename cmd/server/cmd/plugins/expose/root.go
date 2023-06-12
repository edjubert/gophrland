package expose

import (
	"gophrland/cmd/server/cmd/IPC"
)

const EXPOSED_SPECIAL_WORKSPACE = "exposed_special_workspace"

func root(options ExposeOptions) error {
	name := EXPOSED_SPECIAL_WORKSPACE
	if options.Name != "" {
		name = options.Name
	}
	return IPC.ToggleSpecialWorkspace(name)
}
