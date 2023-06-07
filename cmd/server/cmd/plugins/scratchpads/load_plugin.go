package scratchpads

import (
	"fmt"
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
				fmt.Printf("\n[INFO] - Loading Scratchpads '%s' with options\n", name)

				ref := reflect.ValueOf(option)
				fields := reflect.VisibleFields(reflect.TypeOf(option))
				for _, field := range fields {
					if len(field.Index) == 0 {
						continue
					}

					value := ref.Field(field.Index[0])
					if field.Type.String() == "string" && field.Name == "Command" {
						clientExists := false

						fmt.Println("Option Class", option.Class)
						if option.Class != "" {
							clients, err := getClients()
							if err != nil {
								return err
							}

							for _, client := range clients {
								if client.InitialClass == option.Class && !clientExists {
									clientExists = true
									byName[name] = Scratchpad{
										Pid:     client.Pid,
										Options: option,
									}
								}
							}
						}

						fmt.Println("clientExists", clientExists)
						if !clientExists {
							values := strings.Fields(value.String())
							cmd := exec.Command(values[0], values[1:]...)
							if err := cmd.Start(); err != nil {
								return fmt.Errorf("[ERROR] - Could not start '%s' -> %w\n", field.Name, err)
							}

							pid := cmd.Process.Pid
							fmt.Printf("[INFO] - New PID '%s' -> %d\n", name, pid)
							byName[name] = Scratchpad{
								Pid:     pid,
								Options: option,
							}
						}

						fmt.Println("byName ->", byName[name])
					}
				}
			}
		}
	}
	return nil
}
