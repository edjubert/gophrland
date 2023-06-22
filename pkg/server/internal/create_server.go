package internal

import (
	"fmt"
	"log"
	"net"
)

const ServerType = "tcp"

func CreateSocket(signature string) net.Listener {
	socket := "/tmp/hypr/" + signature + "/.gophrland.sock"

	server, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal("create socket: ", err.Error())
	}
	fmt.Println("server socket created")

	return server
}
