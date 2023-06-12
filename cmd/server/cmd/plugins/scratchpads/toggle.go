package scratchpads

import (
	"fmt"
	"gophrland/cmd/server/cmd/IPC"
	"gophrland/cmd/server/cmd/plugins/scratchpads/cmd"
	"time"
)

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

	if err := IPC.MoveToCurrent(monitor.ActiveWorkspace.Id, client.Address); err != nil {
		return err
	}
	if err := IPC.FocusWindow(client.Address); err != nil {
		return err
	}

	return cmd.FromAnimation(client, monitor, animationOptions)
}

func hideClient(client IPC.HyprlandClient, monitor IPC.HyprlandMonitor, animationsOptions cmd.AnimationsOptions) error {
	if err := cmd.ToAnimation(client, monitor, animationsOptions); err != nil {
		fmt.Printf("[ERROR] - %v\n", err)
		return err
	}

	time.Sleep(time.Millisecond * 200)
	return IPC.MoveToSpecialNamed(client.Address)
}