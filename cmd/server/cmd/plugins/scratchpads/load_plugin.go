package scratchpads

import (
	"fmt"
	"reflect"
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
				fmt.Printf("\nLoading Scratchpads '%s' with options\n", name)

				ref := reflect.ValueOf(option)
				fields := reflect.VisibleFields(reflect.TypeOf(option))
				for _, field := range fields {
					fmt.Printf("%d - Key: %s\tType: %s\tValue: %v\n", field.Index[0], field.Name, field.Type, ref.Field(field.Index[0]))
				}
			}
		}
	}
}
