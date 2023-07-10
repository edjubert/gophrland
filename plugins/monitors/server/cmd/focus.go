package cmd

import (
	"fmt"
	IPC "github.com/edjubert/hyprland-ipc-go"
)

const (
	Next     = "next"
	Previous = "previous"
)

func getMonitorIndex(activeMonitor IPC.HyprlandMonitor, monitors []IPC.HyprlandMonitor) int {
	for idx, monitor := range monitors {
		if monitor.Id == activeMonitor.Id {
			return idx
		}
	}

	return -1
}

func getNextMonitor(activeMonitorIndex int, monitors []IPC.HyprlandMonitor) IPC.HyprlandMonitor {
	if activeMonitorIndex+1 < len(monitors) {
		return monitors[activeMonitorIndex+1]
	} else {
		return monitors[0]
	}
}

func getPrevMonitor(activeMonitorIndex int, monitors []IPC.HyprlandMonitor) IPC.HyprlandMonitor {
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

	nextMonitor := IPC.HyprlandMonitor{}
	if args[0] == Next {
		nextMonitor = getNextMonitor(activeMonitorIndex, monitors)
	} else if args[0] == Previous {
		nextMonitor = getPrevMonitor(activeMonitorIndex, monitors)
	}

	return IPC.FocusMonitor(nextMonitor)
}
