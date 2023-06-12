package scratchpads

import (
	"encoding/json"
	"fmt"
	"gophrland/cmd/server/cmd/IPC"
	"gophrland/cmd/server/cmd/plugins/scratchpads/cmd"
	"os/exec"
	"time"
)

func getActiveClient() (IPC.HyprlandClient, error) {
	clientJSON, err := exec.Command("hyprctl", "activewindow", "-j").Output()
	if err != nil {
		return IPC.HyprlandClient{}, err
	}

	var activewindow IPC.HyprlandClient
	if err := json.Unmarshal(clientJSON, &activewindow); err != nil {
		return IPC.HyprlandClient{}, err
	}

	return activewindow, nil
}

func getClients() ([]IPC.HyprlandClient, error) {
	clientsJSON, err := exec.Command("hyprctl", "clients", "-j").Output()
	if err != nil {
		return nil, fmt.Errorf("[ERROR] - Cannot execute command -> %w\n", err)
	}

	var clients []IPC.HyprlandClient
	if err := json.Unmarshal(clientsJSON, &clients); err != nil {
		return nil, fmt.Errorf("[ERROR] - Cannot unmarshall clients -> %w\n", err)
	}

	return clients, nil
}

func getClientByPID(clients []IPC.HyprlandClient, pid int) (IPC.HyprlandClient, error) {
	for _, client := range clients {
		if client.Pid == pid {
			return client, nil
		}
	}

	return IPC.HyprlandClient{}, fmt.Errorf("could not found client")
}

func getClientByClassName(clients []IPC.HyprlandClient, class string) (IPC.HyprlandClient, error) {
	for _, client := range clients {
		if client.Class == class {
			return client, nil
		}
	}

	return IPC.HyprlandClient{}, fmt.Errorf("could not found client")
}

func getMonitors() ([]IPC.HyprlandMonitor, error) {
	monitorsJSON, err := exec.Command("hyprctl", "monitors", "-j").Output()
	if err != nil {
		return nil, err
	}

	var monitors []IPC.HyprlandMonitor
	if err := json.Unmarshal(monitorsJSON, &monitors); err != nil {
		return nil, err
	}

	return monitors, nil
}

func getActiveMonitor(monitors []IPC.HyprlandMonitor) (IPC.HyprlandMonitor, error) {
	for _, monitor := range monitors {
		if monitor.Focused {
			return monitor, nil
		}
	}

	return IPC.HyprlandMonitor{}, fmt.Errorf("not found")
}

func getOption(scratchpadName string, options []map[string]ScratchpadOptions) ScratchpadOptions {
	for _, scratchpad := range options {
		for name, option := range scratchpad {
			if scratchpadName == name {
				return option
			}
		}
	}

	return ScratchpadOptions{}
}

func (scratchpad *Scratchpad) updateScratchpad(options ScratchpadOptions) error {
	fmt.Println("[INFO] Updating scratchpad")
	clients, err := getClients()
	if err != nil {
		return err
	}
	client, err := getClientByClassName(clients, options.Class)
	*scratchpad = Scratchpad{
		Pid:     client.Pid,
		Options: options,
	}

	return nil
}

func toggle(args []string, options []map[string]ScratchpadOptions) error {
	if len(args) > 1 {
		return fmt.Errorf("to many arguments\n")
	}

	option := getOption(args[0], options)

	scratchpad := byName[args[0]]
	if scratchpad.Pid == 0 {
		if option.Class == "" {
			return fmt.Errorf("could not find scratchpad for %s\n", args[0])
		}

		if err := scratchpad.updateScratchpad(option); err != nil {
			return err
		}
	}

	monitors, err := getMonitors()
	if err != nil {
		return err
	}

	activeMonitor, err := getActiveMonitor(monitors)
	if err != nil {
		return err
	}

	clients, err := getClients()
	if err != nil {
		return err
	}

	client, err := getClientByPID(clients, scratchpad.Pid)
	fmt.Printf("Found the client: %d\n%v\n", scratchpad.Pid, client)

	opts := cmd.AnimationsOptions{
		Margin:    scratchpad.Options.Margin,
		Animation: scratchpad.Options.Animation,
	}

	if client.Workspace.Id < 0 {
		if err := showClient(client, activeMonitor, opts); err != nil {
			return err
		}
	} else {
		if err := hideClient(client, activeMonitor, opts); err != nil {
			return err
		}
	}
	return nil
}

func showClient(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, animationOptions cmd.AnimationsOptions) error {
	if err := cmd.ToAnimation(client, monitor, animationOptions); err != nil {
		return err
	}

	moveToCurrent := exec.Command("hyprctl", "dispatch", "movetoworkspacesilent", fmt.Sprintf("%d,address:%s", monitor.ActiveWorkspace.Id, client.Address))
	if err := moveToCurrent.Run(); err != nil {
		return err
	}

	focus := exec.Command("hyprctl", "dispatch", "focuswindow", fmt.Sprintf("address:%s", client.Address))
	if err := focus.Run(); err != nil {
		return err
	}

	return cmd.FromAnimation(client, monitor, animationOptions)
}

func hideClient(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, animationsOptions cmd.AnimationsOptions) error {
	fmt.Println("Hiding client")
	if err := cmd.ToAnimation(client, monitor, animationsOptions); err != nil {
		fmt.Printf("[ERROR] - %v\n", err)
		return err
	}

	fmt.Println("hide client")
	time.Sleep(time.Millisecond * 200)
	return exec.Command("hyprctl", "dispatch", "movetoworkspacesilent", fmt.Sprintf("special:scratchpads,address:%s", client.Address)).Run()
}
