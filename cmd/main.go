// socket-server project main.go
package main

import (
	"gophrland/cmd/client"
	"gophrland/cmd/server"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		server.Handle()
	} else {
		client.Handle()
	}
}
