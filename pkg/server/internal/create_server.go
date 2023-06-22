package internal

import (
	"fmt"
	"github.com/edjubert/gophrland/pkg/client/pkg/tools"
	"log"
	"net"
)

const ServerType = "tcp"

func CreateSocket() net.Listener {
	signature := tools.GetSignature()
	socket := "/tmp/hypr/" + signature + "/.gophrland.sock"

	server, err := net.Listen("unix", socket)
	if err != nil {
		log.Fatal("create socket: ", err.Error())
	}
	fmt.Println("server socket created")

	return server
}
