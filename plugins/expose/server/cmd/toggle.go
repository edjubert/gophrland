package cmd

import (
	"github.com/edjubert/hyprland-ipc-go/hyprctl/dispatch"
	"github.com/edjubert/hyprland-ipc-go/hyprctl/get"
)

func toggle(options ExposeOptions) error {
	name := EXPOSED_SPECIAL_WORKSPACE
	if options.Name != "" {
		name = options.Name
	}

	getter := get.Get{}
	client, err := getter.ActiveClient()
	if err != nil {
		return err
	}

	move := dispatch.Move{}
	if client.Workspace.Id < 0 {
		if err := move.ClientToCurrent(client.Address); err != nil {
			return err
		}
	} else {
		if err := move.ToSpecialNamed(name, client.Address); err != nil {
			return err
		}
	}

	return nil
}
