// socket-server project main.go
package main

import (
	"fmt"
	"gophrland/cmd/internal/server"
)

func main() {
	server.ReadConfig(server.DEFAULT_CONFIG)

	fmt.Println("Server Running...")
	server.Handle()
}
