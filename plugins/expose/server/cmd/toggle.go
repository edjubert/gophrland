package cmd

import (
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
)

func toggle(options ExposeOptions) error {
	name := EXPOSED_SPECIAL_WORKSPACE
	if options.Name != "" {
		name = options.Name
	}

	getter := hyprctl.Get{}
	client, err := getter.ActiveClient()
	if err != nil {
		return err
	}

	dispatch := hyprctl.Dispatch{}
	if client.Workspace.Id < 0 {
		if err := dispatch.Move.ClientToCurrent(client.Address); err != nil {
			return err
		}
	} else {
		if err := dispatch.Move.ToSpecialNamed(name, client.Address); err != nil {
			return err
		}
	}

	return nil
}
