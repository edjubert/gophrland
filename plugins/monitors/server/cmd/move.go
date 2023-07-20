package cmd

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
	"github.com/edjubert/hyprland-ipc-go/types"
)

func move(args []string, opts MonitorsOptions) error {
	if len(args) > 1 {
		return fmt.Errorf("[ERROR] - too many arguments\n")
	}

	getter := hyprctl.Get{}
	monitors, err := getter.Monitors("-j")
	if err != nil {
		return err
	}

	if len(monitors) == 1 {
		return nil
	}

	activeMonitor, err := getter.ActiveMonitor(monitors)
	if err != nil {
		return err
	}

	activeMonitorIndex := getMonitorIndex(activeMonitor, monitors)
	if activeMonitorIndex == -1 {
		return fmt.Errorf("[ERROR] - Could not find monitor index")
	}

	activeClient, err := getter.ActiveClient()
	if err != nil {
		return err
	}

	nextMonitor := types.HyprlandMonitor{}
	if args[0] == Next {
		nextMonitor = getNextMonitor(activeMonitorIndex, monitors)
	} else if args[0] == Previous {
		nextMonitor = getPrevMonitor(activeMonitorIndex, monitors)
		fmt.Println("next monitor", nextMonitor)
	}

	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.ToWorkspaceName(nextMonitor.ActiveWorkspace.Name, activeClient.Address)
}
