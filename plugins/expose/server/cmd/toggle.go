package cmd

import (
	IPC "github.com/edjubert/hyprland-ipc-go"
)

func toggle(options ExposeOptions) error {
	name := EXPOSED_SPECIAL_WORKSPACE
	if options.Name != "" {
		name = options.Name
	}
	client, err := IPC.GetActiveClient()
	if err != nil {
		return err
	}

	if client.Workspace.Id < 0 {
		if err := IPC.MoveToCurrent(client.Address); err != nil {
			return err
		}
	} else {
		if err := IPC.MoveToSpecialNamed(name, client.Address); err != nil {
			return err
		}
	}

	return nil
}
