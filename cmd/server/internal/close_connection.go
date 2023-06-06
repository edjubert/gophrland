package internal

import (
	"net"
)

func closeConnection(conn net.Conn) {
	if err := conn.Close(); err != nil {
		panic(err)
	}
}
