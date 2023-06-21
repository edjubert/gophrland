package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/pkg/server/pkg/IPC"
)

const EXPOSED_SPECIAL_WORKSPACE = "exposed_special_workspace"

func show(options ExposeOptions) error {
	fmt.Println("Hello spechial")
	name := EXPOSED_SPECIAL_WORKSPACE
	if options.Name != "" {
		name = options.Name
	}
	return IPC.ToggleSpecialWorkspace(name)
}
