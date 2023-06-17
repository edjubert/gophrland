// socket-server project main.go
package main

import (
	cmd2 "gophrland/pkg/client/cmd"
	"gophrland/pkg/server/cmd"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		cmd.New()
	} else {
		cmd2.New()
	}
}
