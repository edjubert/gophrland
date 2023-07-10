package cmd

import (
	"fmt"
	IPC "github.com/edjubert/hyprland-ipc-go"
)

func move(args []string, opts MonitorsOptions) error {
	if len(args) > 1 {
		return fmt.Errorf("[ERROR] - too many arguments\n")
	}

	monitors, err := IPC.Monitors("-j")
	if err != nil {
		return err
	}

	if len(monitors) == 1 {
		return nil
	}

	activeMonitor, err := IPC.ActiveMonitor(monitors)
	if err != nil {
		return err
	}

	activeMonitorIndex := getMonitorIndex(activeMonitor, monitors)
	if activeMonitorIndex == -1 {
		return fmt.Errorf("[ERROR] - Could not find monitor index")
	}

	activeClient, err := IPC.GetActiveClient()
	if err != nil {
		return err
	}

	nextMonitor := IPC.HyprlandMonitor{}
	if args[0] == Next {
		nextMonitor = getNextMonitor(activeMonitorIndex, monitors)
	} else if args[0] == Previous {
		nextMonitor = getPrevMonitor(activeMonitorIndex, monitors)
		fmt.Println("next monitor", nextMonitor)
	}

	return IPC.MoveToWorkspace(nextMonitor.ActiveWorkspace.Name, activeClient.Address)
}
