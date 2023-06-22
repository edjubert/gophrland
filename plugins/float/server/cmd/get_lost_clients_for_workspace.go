package cmd

import (
	IPC "github.com/edjubert/hyprland-ipc-go"
)

func getLostClientsForWorkspace(workspace IPC.HyprlandWorkspace, monitor IPC.HyprlandMonitor, opts BringFloatOptions) ([]LostClient, error) {
	floatingClients, err := IPC.GetWorkspaceFloatingClients(workspace)
	if err != nil {
		return nil, err
	}

	var lostClients []LostClient
	for _, client := range floatingClients {
		lostClient, err := clientIsLost(client, monitor, opts)
		if err != nil {
			return nil, err
		}

		if lostClient.Top || lostClient.Bottom || lostClient.Left || lostClient.Right {
			lostClients = append(lostClients, lostClient)
		}
	}

	return lostClients, nil
}
