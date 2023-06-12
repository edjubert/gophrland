package scratchpads

import (
	"fmt"
	"gophrland/cmd/server/cmd/IPC"
	cmd2 "gophrland/cmd/server/cmd/plugins/scratchpads/cmd"
	"os/exec"
	"reflect"
	"strings"
)

type ScratchpadOptions struct {
	Command   string `yaml:"command"`
	Animation string `yaml:"animation,omitempty"`
	Unfocus   string `yaml:"unfocus,omitempty"`
	Margin    int    `yaml:"margin,omitempty"`
	Class     string `yaml:"class,omitempty"`
}

type Scratchpad struct {
	Pid     int
	Options ScratchpadOptions
}

var byName = make(map[string]Scratchpad)

func LoadPlugin(options []map[string]ScratchpadOptions) error {
	if len(options) > 0 {
		for _, scratchpad := range options {
			for name, option := range scratchpad {
				fmt.Printf("[INFO] - Loading Scratchpads '%s' with options\n", name)

				ref := reflect.ValueOf(option)
				fields := reflect.VisibleFields(reflect.TypeOf(option))
				for _, field := range fields {
					if len(field.Index) == 0 {
						continue
					}

					value := ref.Field(field.Index[0])
					if field.Type.String() == "string" && field.Name == "Command" {
						if option.Class == "" {
							values := strings.Fields(value.String())
							cmd := exec.Command(values[0], values[1:]...)
							if err := cmd.Start(); err != nil {
								return fmt.Errorf("[ERROR] - Could not start '%s' -> %w\n", field.Name, err)
							}

							pid := cmd.Process.Pid

							clients, err := IPC.GetClients()
							if err != nil {
							}
							client, err := IPC.GetClientByPID(clients, pid)
							if err != nil {
							}

							monitors, err := IPC.Monitors("-j")
							monitor, err := IPC.ActiveMonitor(monitors)

							opts := cmd2.AnimationsOptions{
								Margin:    option.Margin,
								Animation: option.Animation,
							}
							if err := cmd2.ToAnimation(client, monitor, opts); err != nil {

							}

							byName[name] = Scratchpad{
								Pid:     pid,
								Options: option,
							}
						}
					}
				}
			}
		}
	}
	return nil
}
