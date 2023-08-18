package cmd

import (
	"fmt"
	server "github.com/edjubert/gophrland/plugins/scratchpads/server/pkg"
	"github.com/edjubert/hyprland-ipc-go/hyprctl"
	"os/exec"
	"reflect"
	"strings"
)

type ScratchpadFloatOptions struct {
	Animation string `yaml:"animation,omitempty"`
	Margin    int    `yaml:"margin,omitempty"`
	Width     string `yaml:"width,omitempty"`
	Height    string `yaml:"height,omitempty"`
}

type ScratchpadOptions struct {
	Command      string                 `yaml:"command"`
	Float        bool                   `yaml:"float,omitempty"`
	FloatOptions ScratchpadFloatOptions `yaml:"floatOpts,omitempty"`
	Unfocus      string                 `yaml:"unfocus,omitempty"`
	Class        string                 `yaml:"class,omitempty"`
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
				fmt.Printf("[INFO] - Loading Scratchpad '%s' with options\n", name)

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
							if len(values) == 0 {
								return fmt.Errorf("[ERROR] - No values")
							}
							cmd := exec.Command(values[0], values[1:]...)
							if err := cmd.Start(); err != nil {
								return fmt.Errorf("[ERROR] - Could not start '%s' -> %w\n", field.Name, err)
							}

							pid := cmd.Process.Pid

							getter := hyprctl.Get{}
							clients, err := getter.Clients()
							if err != nil {
							}
							client, err := getter.ClientByPID(clients, pid)
							if err != nil {
							}

							monitors, err := getter.Monitors("-j")
							monitor, err := getter.ActiveMonitor(monitors)

							if option.Float {
								opts := server.AnimationsOptions{
									Margin:    option.FloatOptions.Margin,
									Animation: option.FloatOptions.Animation,
									Width:     option.FloatOptions.Width,
									Height:    option.FloatOptions.Height,
								}
								if err := server.ResizeClient(client, opts.Width, opts.Height); err != nil {
								}
								if err := server.ToAnimation(client, monitor, opts); err != nil {
								}
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
