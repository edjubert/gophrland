package IPC

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os/exec"
	"strconv"
)

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

func GetWorkspaceFloatingClients(workspace HyprlandWorkspace) ([]HyprlandClient, error) {
	clients, err := GetClients()
	if err != nil {
		return nil, err
	}

	var workspaceClients []HyprlandClient
	for _, client := range clients {
		if client.Workspace.Id == workspace.Id {
			workspaceClients = append(workspaceClients, client)
		}
	}

	return workspaceClients, nil
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

	return HyprlandClient{}, fmt.Errorf("[ERROR] - could not found client")
}

func GetClientByClassName(clients []HyprlandClient, class string) (HyprlandClient, error) {
	for _, client := range clients {
		if client.Class == class {
			return client, nil
		}
	}

	return HyprlandClient{}, fmt.Errorf("[ERROR] - could not found client")
}

func Monitors(format string) ([]HyprlandMonitor, error) {
	if format != "" && format != "-j" {
		return nil, fmt.Errorf("[ERROR] - wrong monitor formats")
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

	return HyprlandMonitor{}, fmt.Errorf("[ERROR] - not found")
}

func SendNotification(time int, msgType, msg string) error {
	icon := -1
	prefix := "  [Gophrland]"

	switch msgType {
	case "warning":
		icon = 0
	case "info":
		icon = 1
	case "notice":
		icon = 2
	case "error":
		icon = 3
	case "question":
		icon = 4
	case "checkmark":
		icon = 5
	default:
		icon = -1
	}

	color := "rgb(ff1ea3)"
	return runHyprctlCmd(fmt.Sprintf("notify %d %d %s %s: %s", icon, time, color, prefix, msg))
}

func runHyprctlCmd(cmd string) error {
	conn, err := ConnectHyprctl()
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	if _, err := conn.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}

func MoveWindowPixelExact(x, y int, address string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch movewindowpixel exact %d %d,address:%s", x, y, address))
}

func ToggleSpecialWorkspace(name string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch togglespecialworkspace %s", name))
}

func CenterFloatingClient(client HyprlandClient, monitor HyprlandMonitor) error {
	margin := 100
	randFactorX := client.Size[0]
	randFactorY := client.Size[1]
	randX := rand.Intn(randFactorX)
	randY := rand.Intn(randFactorY)
	centerX := (monitor.X + monitor.Width - monitor.Width/2) - client.Size[0]/2 - randFactorX/2 + randX
	centerY := (monitor.Y + monitor.Height - monitor.Height/2) - client.Size[1]/2 - randFactorY/2 + randY + margin

	return runHyprctlCmd(fmt.Sprintf("dispatch movewindowpixel exact %d %d,address:%s", centerX, centerY, client.Address))
}

func MoveToCurrent(address string) error {
	monitors, err := Monitors("-j")
	if err != nil {
		return err
	}

	monitor, err := ActiveMonitor(monitors)
	if err != nil {
		return err
	}

	if err := MoveToWorkspaceID(monitor.ActiveWorkspace.Id, address); err != nil {
		return err
	}

	return nil
}

func FocusCurrentWorkspace(currentWorkspaceId int) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch workspace %d", currentWorkspaceId))
}

func MoveToWorkspaceID(currentWorkspaceID int, address string) error {
	return MoveToWorkspaceSilent(strconv.Itoa(currentWorkspaceID), address)
}
func FocusWindow(address string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch focuswindow address,%s", address))
}

func MoveToWorkspaceSilent(name, address string) error {
	return runHyprctlCmd(fmt.Sprintf("dispatch movetoworkspacesilent %s,address:%s", name, address))
}

func GetActiveWorkspace() (HyprlandWorkspace, error) {
	activeClient, err := GetActiveClient()
	if err != nil {
		return HyprlandWorkspace{}, err
	}

	return activeClient.Workspace, nil
}

func GetWorkspaces() ([]HyprlandWorkspace, error) {
	ret, err := exec.Command("hyprctl", "workspaces", "-j").Output()
	if err != nil {
		return nil, err
	}

	var workspaces []HyprlandWorkspace
	if err := json.Unmarshal(ret, &workspaces); err != nil {
		return nil, err
	}

	fmt.Println(workspaces)
	return workspaces, nil
}
func MoveToSpecialNamed(specialName, address string) error {
	if specialName != "" {
		specialName = fmt.Sprintf(":%s", specialName)
	}
	return MoveToWorkspaceSilent("special"+specialName, address)
}
