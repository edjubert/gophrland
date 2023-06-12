package internal

import (
	"fmt"
	"gophrland/cmd/server/cmd/config"
	"gophrland/cmd/server/cmd/plugins/expose"
	"gophrland/cmd/server/cmd/plugins/scratchpads"
	"net"
	"strings"
)

func callCommand(command string, opts config.Options) error {
	fields := strings.Fields(command)
	plugin := fields[0]

	fmt.Println("len: ", len(fields), fields[1:])

	switch plugin {
	case config.SCRATCHPADS:
		cmd := fields[1]
		args := fields[2:]
		return scratchpads.Command(cmd, args, opts.Scratchpads)
	case config.EXPOSE:
		cmd := ""
		if len(fields[1:]) > 0 {
			cmd = fields[1]
		}
		return expose.Command(cmd, opts.Expose)
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
	if err := callCommand(string(buffer[:mLen]), loadedConf.Options); err != nil {
		fmt.Printf("%v\n", err)
	}

	if _, err := connection.Write([]byte("ok")); err != nil {
		fmt.Printf("Error writing response: %v\n", err)
		return
	}
}
