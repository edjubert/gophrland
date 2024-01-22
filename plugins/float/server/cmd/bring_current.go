package cmd

import (
	"github.com/edjubert/hyprland-ipc-go/hyprctl/dispatch"
	"github.com/edjubert/hyprland-ipc-go/hyprctl/get"
	"github.com/edjubert/hyprland-ipc-go/types"
)

const OFFSET = 0.5

type LostClient struct {
	Client                   types.HyprlandClient
	Left, Right, Top, Bottom bool
}

func bringCurrent(opts FloatOptions) error {
	getter := get.Get{}
	activeWorkspace, err := getter.ActiveWorkspace()
	if err != nil {
		return err
	}

	monitors, err := getter.Monitors("-j")
	if err != nil {
		return err
	}

	monitor, err := getter.ActiveMonitor(monitors)
	if err != nil {
		return err
	}
	lostClients, err := getLostClientsForWorkspace(activeWorkspace, monitor, opts)
	if err != nil {
		return err
	}

	move := dispatch.Move{}
	for _, client := range lostClients {
		if err := move.CenterFloatingClient(client.Client, monitor, opts.RandomizeCenter); err != nil {
			return err
		}
	}

	return nil
}
