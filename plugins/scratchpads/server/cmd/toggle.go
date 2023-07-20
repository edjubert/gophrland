package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/plugins/scratchpads/server/pkg"
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
	"github.com/edjubert/hyprland-ipc-go/types"
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
	getter := hyprctl.Get{}
	clients, err := getter.Clients()
	if err != nil {
		return err
	}
	client, err := getter.ClientByClassName(clients, options.Class)
	*scratchpad = Scratchpad{
		Pid:     client.Pid,
		Options: options,
	}

	dispatch := hyprctl.Dispatch{}
	if !options.Float && client.Floating {
		if err := dispatch.Toggle.Floating(client.Address); err != nil {
			return err
		}
	}

	return nil
}

var blockListener = false

func toggle(args []string, options []map[string]ScratchpadOptions) error {
	blockListener = true
	if len(args) > 1 {
		return fmt.Errorf("[ERROR] - too many arguments\n")
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

	getter := hyprctl.Get{}
	monitors, err := getter.Monitors("-j")
	if err != nil {
		return err
	}

	activeMonitor, err := getter.ActiveMonitor(monitors)
	if err != nil {
		return err
	}

	clients, err := getter.Clients()
	if err != nil {
		return err
	}

	client, err := getter.ClientByPID(clients, scratchpad.Pid)

	if scratchpad.Options.Float || client.Floating {
		opts := pkg.AnimationsOptions{
			Margin:    scratchpad.Options.FloatOptions.Margin,
			Animation: scratchpad.Options.FloatOptions.Animation,
			Width:     scratchpad.Options.FloatOptions.Width,
			Height:    scratchpad.Options.FloatOptions.Height,
		}

		if client.Workspace.Id < 0 {
			if err := showFloatingClient(client, activeMonitor, opts); err != nil {
				return err
			}
		} else {
			if err := hideFloatingClient(client, activeMonitor, opts); err != nil {
				return err
			}
		}
	} else {
		if client.Workspace.Id < 0 {
			if err := showClient(client, activeMonitor); err != nil {
				return err
			}
		} else {
			if err := hideClient(client); err != nil {
				return err
			}
		}
	}

	time.Sleep(time.Second * 2)
	blockListener = false
	return nil
}

func showClient(client types.HyprlandClient, monitor types.HyprlandMonitor) error {
	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.ToWorkspaceName(monitor.ActiveWorkspace.Name, client.Address)
}
func hideClient(client types.HyprlandClient) error {
	dispatch := hyprctl.Dispatch{}
	return dispatch.Move.ToSpecialNamed(SCRATCHPADS_SPECIAL_WORKSPACE, client.Address)
}

func showFloatingClient(client types.HyprlandClient, monitor types.HyprlandMonitor, animationOptions pkg.AnimationsOptions) error {
	if err := pkg.ToAnimation(client, monitor, animationOptions); err != nil {
		return err
	}

	dispatch := hyprctl.Dispatch{}
	if err := dispatch.Move.ToWorkspaceName(monitor.ActiveWorkspace.Name, client.Address); err != nil {
		return err
	}

	return pkg.FromAnimation(client, monitor, animationOptions)
}

func hideFloatingClient(client types.HyprlandClient, monitor types.HyprlandMonitor, animationsOptions pkg.AnimationsOptions) error {
	if err := pkg.ToAnimation(client, monitor, animationsOptions); err != nil {
		fmt.Printf("[ERROR] - %v\n", err)
		return err
	}

	dispatch := hyprctl.Dispatch{}
	time.Sleep(time.Millisecond * 200)
	return dispatch.Move.ToSpecialNamed(SCRATCHPADS_SPECIAL_WORKSPACE, client.Address)
}
