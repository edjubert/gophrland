package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/plugins/scratchpads/server/pkg"
	IPC "github.com/edjubert/hyprland-ipc-go"
)

type ScratchpadsAndClients struct {
	Scratchpad Scratchpad
	Client     IPC.HyprlandClient
}

const (
	HideOption = "hide"
)

func getAllScratchpadsAndClients(options []map[string]ScratchpadOptions) ([]ScratchpadsAndClients, error) {
	clients, err := IPC.GetClients()
	if err != nil {
		return nil, err
	}

	var scratchpads []ScratchpadsAndClients
	for _, option := range options {
		for name, opts := range option {
			if byName[name].Pid == 0 {
				client, err := IPC.GetClientByClassName(clients, opts.Class)
				if err != nil {
					fmt.Println("[ERR] 33 - ", err)
				}

				scratchpads = append(scratchpads, ScratchpadsAndClients{
					Scratchpad: Scratchpad{
						Pid:     client.Pid,
						Options: opts,
					},
					Client: client,
				})
				continue
			}

			for _, c := range clients {
				if c.Pid == byName[name].Pid {
					scratchpads = append(scratchpads, ScratchpadsAndClients{
						Scratchpad: byName[name],
						Client:     c,
					})
					break
				}
			}
		}
	}

	return scratchpads, nil
}

func hideOnUnfocused(clientAddresses []string, options []map[string]ScratchpadOptions) error {
	scratches, err := getAllScratchpadsAndClients(options)
	if err != nil {
		return err
	}

	currentWorkspace, err := IPC.GetActiveWorkspace()
	if err != nil {
		return err
	}

	var toKeepScratches []ScratchpadsAndClients
	for _, client := range clientAddresses {
		for _, scratch := range scratches {
			if scratch.Client.Address == "0x"+client && scratch.Client.Workspace.Id == currentWorkspace.Id {
				toKeepScratches = append(toKeepScratches, scratch)
			}
		}
	}
	monitors, err := IPC.Monitors("-j")
	activeMonitor, err := IPC.ActiveMonitor(monitors)

	for _, scratch := range scratches {
		doKeep := false
		for _, toKeep := range toKeepScratches {
			if toKeep.Client.Pid == scratch.Client.Pid && scratch.Client.Workspace.Id == currentWorkspace.Id {
				doKeep = true
			}
		}

		if !doKeep && !blockListener && scratch.Scratchpad.Options.Unfocus == HideOption {
			opts := pkg.AnimationsOptions{
				Margin:    scratch.Scratchpad.Options.Margin,
				Animation: scratch.Scratchpad.Options.Animation,
			}

			fmt.Println("Hidding ", scratch.Client.Title)
			if err := hideClient(scratch.Client, activeMonitor, opts); err != nil {
				fmt.Println("[ERROR]", err)
			}
		}
	}

	return nil
}

func showOnUrgent(clientAddresses []string, options []map[string]ScratchpadOptions) error {
	scratches, err := getAllScratchpadsAndClients(options)
	if err != nil {
		return err
	}

	currentWorkspace, err := IPC.GetActiveWorkspace()
	if err != nil {
		return err
	}

	var toKeepScratches []ScratchpadsAndClients
	for _, client := range clientAddresses {
		for _, scratch := range scratches {
			if scratch.Client.Address == "0x"+client {
				toKeepScratches = append(toKeepScratches, scratch)
			}
		}
	}
	monitors, err := IPC.Monitors("-j")
	activeMonitor, err := IPC.ActiveMonitor(monitors)

	for _, scratch := range scratches {
		for _, toKeep := range toKeepScratches {
			if toKeep.Client.Pid == scratch.Client.Pid && scratch.Client.Workspace.Id == currentWorkspace.Id && !blockListener {
				opts := pkg.AnimationsOptions{
					Margin:    scratch.Scratchpad.Options.Margin,
					Animation: scratch.Scratchpad.Options.Animation,
				}

				if err := showClient(scratch.Client, activeMonitor, opts); err != nil {
					fmt.Println("[ERROR]", err)
				}
			}
		}

	}

	return nil
}

func parseEvents(options []map[string]ScratchpadOptions) func(socketMessages IPC.HyprSocketMessage) {
	return func(socketMessages IPC.HyprSocketMessage) {
		if len(socketMessages[IPC.ActiveWindowV2]) > 0 {
			if err := hideOnUnfocused(socketMessages[IPC.ActiveWindowV2], options); err != nil {
				fmt.Printf("[ERROR] - %v\n", err)
			}
		}

		if len(socketMessages[IPC.Urgent]) > 0 {
			if err := showOnUrgent(socketMessages[IPC.Urgent], options); err != nil {
				fmt.Printf("[ERROR] - %v\n", err)
			}
		}
	}
}

func LoadEventCallbacks(options []map[string]ScratchpadOptions) {
	callbacks := []IPC.HyprlandCallback{
		parseEvents(options),
	}

	fmt.Println("Loading scratchpad callbacks", len(callbacks))
	go IPC.ConnectEvents(callbacks)
}
