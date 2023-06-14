package client

import (
	"gophrland/cmd/client/cmd"
	"net"
)

const (
	ServerHost = "localhost"
	ServerPort = "9988"
	ServerType = "tcp"
)

func startConnection() net.Conn {
	//establish connection
	connection, err := net.Dial(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		panic(err)
	}

	return connection
}

func New() {
	connection := startConnection()
	cmd.InitRootCmd(connection)

	if err := connection.Close(); err != nil {
		panic(err)
	}
}
