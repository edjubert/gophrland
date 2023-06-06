package scratchpads

import (
	"fmt"
	"os/exec"
	"reflect"
	"strings"
)

type Scratchpad struct {
	Command   string `yaml:"command"`
	Animation string `yaml:"animation"`
	Unfocus   string `yaml:"unfocus"`
}

func LoadPlugin(options []map[string]Scratchpad) {
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
						values := strings.Fields(value.String())

						cmd := exec.Command(values[0], values[1:]...)
						if err := cmd.Start(); err != nil {
							fmt.Printf("[ERROR] - Could not start '%s' -> %v\n", field.Name, err)
						}

						pid := cmd.Process.Pid
						fmt.Println("[INFO] - New PID -> ", pid)
					}
				}
			}
		}
	}
}
