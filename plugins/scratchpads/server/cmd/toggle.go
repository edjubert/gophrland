package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/plugins/scratchpads/server/pkg"
	IPC "github.com/edjubert/hyprland-ipc-go"
	"time"
)

const SCRATCHPADS_SPECIAL_WORKSPACE = "scratchpads_special_workspace"

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
	clients, err := IPC.GetClients()
	if err != nil {
		return err
	}
	client, err := IPC.GetClientByClassName(clients, options.Class)
	*scratchpad = Scratchpad{
		Pid:     client.Pid,
		Options: options,
	}

	return nil
}

func toggle(args []string, options []map[string]ScratchpadOptions) error {
	if len(args) > 1 {
		return fmt.Errorf("[ERROR] - to many arguments\n")
	}

	option := getOption(args[0], options)

	scratchpad := byName[args[0]]
	if scratchpad.Pid == 0 {
		if option.Class == "" {
			return fmt.Errorf("[ERROR] - could not find scratchpad for %s\n", args[0])
		}

		if err := scratchpad.updateScratchpad(option); err != nil {
			return err
		}
	}

	monitors, err := IPC.Monitors("-j")
	if err != nil {
		return err
	}

	activeMonitor, err := IPC.ActiveMonitor(monitors)
	if err != nil {
		return err
	}

	clients, err := IPC.GetClients()
	if err != nil {
		return err
	}

	client, err := IPC.GetClientByPID(clients, scratchpad.Pid)

	opts := pkg.AnimationsOptions{
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

func showClient(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, animationOptions pkg.AnimationsOptions) error {
	if err := pkg.ToAnimation(client, monitor, animationOptions); err != nil {
		return err
	}

	if err := IPC.MoveToWorkspaceID(monitor.ActiveWorkspace.Id, client.Address); err != nil {
		return err
	}
	if err := IPC.FocusWindow(client.Address); err != nil {
		return err
	}

	return pkg.FromAnimation(client, monitor, animationOptions)
}

func hideClient(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, animationsOptions pkg.AnimationsOptions) error {
	if err := pkg.ToAnimation(client, monitor, animationsOptions); err != nil {
		fmt.Printf("[ERROR] - %v\n", err)
		return err
	}

	time.Sleep(time.Millisecond * 200)
	return IPC.MoveToSpecialNamed(SCRATCHPADS_SPECIAL_WORKSPACE, client.Address)
}
