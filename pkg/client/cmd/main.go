package cmd

import (
	"github.com/urfave/cli/v2"
	"gophrland/pkg/client/pkg/plugins/bring_float/cmd"
	"gophrland/pkg/client/pkg/plugins/expose"
	cmd2 "gophrland/pkg/client/pkg/plugins/scratchpads/cmd"
	"log"
	"net"
	"os"
)

const (
	ServerHost = "localhost"
	ServerPort = "9988"
	ServerType = "tcp"
)

func run(conn net.Conn) {
	app := &cli.App{
		Name:  "gophrland",
		Usage: "Client and daemon to manage hyprland plugins",
		Commands: []*cli.Command{
			cmd2.GetScratchpadsCommands(conn),
			expose.GetExposeCommands(conn),
			cmd.GetBringFloatCommands(conn),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
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
	run(connection)

	if err := connection.Close(); err != nil {
		panic(err)
	}
}
