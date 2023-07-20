package cmd

import (
	"fmt"
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
	"github.com/edjubert/hyprland-ipc-go/types"
)

const (
	Next     = "next"
	Previous = "previous"
)

func getMonitorIndex(activeMonitor types.HyprlandMonitor, monitors []types.HyprlandMonitor) int {
	for idx, monitor := range monitors {
		if monitor.Id == activeMonitor.Id {
			return idx
		}
	}

	return -1
}

func getNextMonitor(activeMonitorIndex int, monitors []types.HyprlandMonitor) types.HyprlandMonitor {
	if activeMonitorIndex+1 < len(monitors) {
		return monitors[activeMonitorIndex+1]
	} else {
		return monitors[0]
	}
}

func getPrevMonitor(activeMonitorIndex int, monitors []types.HyprlandMonitor) types.HyprlandMonitor {
	fmt.Println("activemonitorindex", activeMonitorIndex, len(monitors))
	if activeMonitorIndex-1 >= 0 {
		return monitors[activeMonitorIndex-1]
	} else {
		return monitors[len(monitors)-1]
	}
}

func focus(args []string, options MonitorsOptions) error {
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

	nextMonitor := types.HyprlandMonitor{}
	if args[0] == Next {
		nextMonitor = getNextMonitor(activeMonitorIndex, monitors)
	} else if args[0] == Previous {
		nextMonitor = getPrevMonitor(activeMonitorIndex, monitors)
	}

	dispatch := hyprctl.Dispatch{}
	return dispatch.Focus.Monitor(nextMonitor)
}
