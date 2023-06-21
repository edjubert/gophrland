package tools

import (
	"fmt"
	"net"
)

func StartConnection(protocol, host string, port int) net.Conn {
	//establish connection
	connection, err := net.Dial(protocol, fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	return connection
}

func StartTCPConnection(host string, port int) net.Conn {
	return StartConnection("tcp", host, port)
}