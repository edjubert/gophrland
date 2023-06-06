package internal

import (
	"fmt"
	"net"
)

func ProcessClient(connection net.Conn) {
	defer closeConnection(connection)

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		if err.Error() != "EOF" {
			fmt.Printf("Could not read buffer: %v\n", err)
		}
		return
	}

	fmt.Println("Received: ", string(buffer[:mLen]))
	if _, err := connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen]))); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}
}
