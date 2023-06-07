package internal

import (
	"fmt"
	"gophrland/cmd/server/cmd/config"
	"gophrland/cmd/server/cmd/plugins/scratchpads"
	"net"
	"strings"
)

func callCommand(command string) error {
	fields := strings.Fields(command)
	plugin := fields[0]
	cmd := fields[1]
	args := fields[2:]

	switch plugin {
	case config.SCRATCHPADS:
		return scratchpads.Command(cmd, args)
	default:
		return fmt.Errorf("%s is not a recognized command\n", plugin)
	}
}

func ProcessClient(connection net.Conn, loadedConf config.Config) {
	defer closeConnection(connection)

	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		if err.Error() != "EOF" {
			fmt.Printf("Could not read buffer: %v\n", err)
		}
		return
	}
	if err := callCommand(string(buffer[:mLen])); err != nil {
		fmt.Printf("%v\n", err)
	}

	if _, err := connection.Write([]byte("ok")); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}
}
