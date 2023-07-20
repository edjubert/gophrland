package cmd

import (
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
	"github.com/edjubert/hyprland-ipc-go/types"
)

func getLostClientsForWorkspace(workspace types.HyprlandWorkspace, monitor types.HyprlandMonitor, opts FloatOptions) ([]LostClient, error) {
	getter := hyprctl.Get{}
	floatingClients, err := getter.WorkspaceFloatingClients(workspace)
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
