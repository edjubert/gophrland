package internal

import (
	"fmt"
	"gophrland/cmd/server/cmd/config"
	"gophrland/cmd/server/cmd/plugins/bring_float"
	"gophrland/cmd/server/cmd/plugins/expose"
	"gophrland/cmd/server/cmd/plugins/scratchpads"
	"net"
	"strings"
)

func callCommand(command string, opts config.Options) error {
	fields := strings.Fields(command)
	plugin := fields[0]

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
	case config.BRING_FLOAT:
		cmd := fields[1]
		return bring_float.Command(cmd, opts.BringFloat)
	default:
		return fmt.Errorf("[ERROR] - %s is not a recognized command\n", plugin)
	}
}

func ProcessClient(connection net.Conn, loadedConf config.Config) {
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
