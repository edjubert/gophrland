package internal

import (
	"fmt"
	"github.com/edjubert/gophrland/plugins"
	"net"
)

func ProcessClient(connection net.Conn, loadedConf plugins.Config) {
	defer closeConnection(connection)

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		if err.Error() != "EOF" {
			fmt.Printf("[ERROR] - Could not read buffer: %v\n", err)
		}
		return
	}

	if err := callCommand(string(buffer[:mLen]), loadedConf.Options); err != nil {
		fmt.Println(err)
	}

	if _, err := connection.Write([]byte("ok")); err != nil {
		fmt.Printf("[ERROR] - Error writing response: %v\n", err)
		return
	}
}
