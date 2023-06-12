package expose

import (
	"gophrland/cmd/server/cmd/IPC"
)

func toggle(options ExposeOptions) error {
	client, err := IPC.GetActiveClient()
	if err != nil {
		return err
	}

	if client.Workspace.Id < 0 {
		if err := IPC.MoveToCurrent(client.Address); err != nil {
			return err
		}
	} else {
		if err := IPC.MoveToSpecialNamed(EXPOSED_SPECIAL_WORKSPACE, client.Address); err != nil {
			return err
		}

		if err := IPC.FocusWindow(client.Address); err != nil {
			return err
		}
	}

	return nil
}
