package expose

import (
	"fmt"
	"reflect"
)

type Expose struct {
}

func LoadPlugin(options []map[string]Expose) {
	if len(options) > 0 {
		for _, expose := range options {
			for name, option := range expose {
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
