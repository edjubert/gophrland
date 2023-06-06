package client

import (
	"gophrland/cmd/client/pkg/plugins/scratchpads"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
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

func startApp(conn net.Conn) {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosion",
		Commands: []*cli.Command{
			scratchpads.GetApp(conn),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Handle() {
	connection := startConnection()
	startApp(connection)

	if err := connection.Close(); err != nil {
		panic(err)
	}
}
