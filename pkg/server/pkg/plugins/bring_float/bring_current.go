package bring_float

import (
	"gophrland/pkg/server/pkg/IPC"
)

const OFFSET = 0.5

type LostClient struct {
	Client                   IPC.HyprlandClient
	Left, Right, Top, Bottom bool
}

func bringCurrent(opts BringFloatOptions) error {
	activeWorkspace, err := IPC.GetActiveWorkspace()
	if err != nil {
		return err
	}

	monitors, err := IPC.Monitors("-j")
	if err != nil {
		return err
	}
	monitor, err := IPC.ActiveMonitor(monitors)
	if err != nil {
		return err
	}
	lostClients, err := getLostClientsForWorkspace(activeWorkspace, monitor, opts)
	if err != nil {
		return err
	}

	for _, client := range lostClients {
		if err := IPC.CenterFloatingClient(client.Client, monitor); err != nil {
			return err
		}
	}

	return nil
}
