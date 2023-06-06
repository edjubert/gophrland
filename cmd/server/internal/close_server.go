package internal

import "net"

func CloseServer(server net.Listener) {
	if err := server.Close(); err != nil {
		panic(err)
	}
}
