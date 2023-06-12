package IPC

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

const SCRATCHPADS_SPECIAL_WORKSPACE = "scratchpads_special_workspace"

func GetActiveClient() (HyprlandClient, error) {
	clientJSON, err := exec.Command("hyprctl", "activewindow", "-j").Output()
	if err != nil {
		return HyprlandClient{}, err
	}

	var activewindow HyprlandClient
	if err := json.Unmarshal(clientJSON, &activewindow); err != nil {
		return HyprlandClient{}, err
	}

	return activewindow, nil
}

func GetClients() ([]HyprlandClient, error) {
	clientsJSON, err := exec.Command("hyprctl", "clients", "-j").Output()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] - Cannot execute command -> %w\n", err)
	}

	var clients []HyprlandClient
	if err := json.Unmarshal(clientsJSON, &clients); err != nil {
		return nil, fmt.Errorf("[ERROR] - Cannot unmarshall clients -> %w\n", err)
	}

	return clients, nil
}

func GetClientByPID(clients []HyprlandClient, pid int) (HyprlandClient, error) {
	for _, client := range clients {
		if client.Pid == pid {
			return client, nil
		}
	}

	return HyprlandClient{}, fmt.Errorf("could not found client")
}

func GetClientByClassName(clients []HyprlandClient, class string) (HyprlandClient, error) {
	for _, client := range clients {
		if client.Class == class {
			return client, nil
		}
	}

	return HyprlandClient{}, fmt.Errorf("could not found client")
}

func Monitors(format string) ([]HyprlandMonitor, error) {
	if format != "" && format != "-j" {
		return nil, fmt.Errorf("wrong monitor formats")
	}

	monitorsJSON, err := exec.Command("hyprctl", "monitors", format).Output()
	if err != nil {
		return nil, err
	}

	var monitors []HyprlandMonitor
	if err := json.Unmarshal(monitorsJSON, &monitors); err != nil {
		return nil, err
	}

	return monitors, nil
}

func ActiveMonitor(monitors []HyprlandMonitor) (HyprlandMonitor, error) {
	for _, monitor := range monitors {
		if monitor.Focused {
			return monitor, nil
		}
	}

	return HyprlandMonitor{}, fmt.Errorf("not found")
}

func MoveWindowPixelExact(x, y int, address string) error {
	return exec.
		Command("hyprctl", "dispatch", "movewindowpixel", "exact", fmt.Sprintf("%d %d,address:%s", x, y, address)).
		Run()
}

func MoveToCurrent(currentWorkspaceID int, address string) error {
	return MoveToWorkspaceSilent(strconv.Itoa(currentWorkspaceID), address)
}
func FocusWindow(address string) error {
	return exec.Command("hyprctl", "dispatch", "focuswindow", "address:"+address).Run()
}

func MoveToWorkspaceSilent(name, address string) error {
	return exec.Command("hyprctl", "dispatch", "movetoworkspacesilent", fmt.Sprintf("%s,address:%s", name, address)).Run()
}
func MoveToSpecialNamed(address string) error {
	return MoveToWorkspaceSilent("special:"+SCRATCHPADS_SPECIAL_WORKSPACE, address)
}
