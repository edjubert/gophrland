package cmd

import (
	"github.com/urfave/cli/v2"
	"gophrland/cmd/client/pkg/plugins/scratchpads"
	"log"
	"net"
	"os"
)

func InitRootCmd(conn net.Conn) {
	app := &cli.App{
		Name:  "gophrland",
		Usage: "Client and daemon to manage hyprland plugins",
		Commands: []*cli.Command{
			scratchpads.GetScratchpadsCommands(conn),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
