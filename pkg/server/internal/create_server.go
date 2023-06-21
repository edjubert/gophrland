package internal

import (
	"fmt"
	"net"
	"os"
)

const ServerType = "tcp"

func CreateServer(host string, port int) net.Listener {
	server, err := net.Listen(ServerType, fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("[ERROR] - Error listening:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("[INFO] - Listening on %s:%d\n", host, port)
	return server
}
