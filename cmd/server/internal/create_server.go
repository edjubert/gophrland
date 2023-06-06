package internal

import (
	"fmt"
	"net"
	"os"
)

const (
	ServerHost = "localhost"
	ServerPort = "9988"
	ServerType = "tcp"
)

func CreateServer() net.Listener {
	server, err := net.Listen(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	fmt.Println("Listening on " + ServerHost + ":" + ServerPort)

	return server
}
