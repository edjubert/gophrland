package bring_float

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
)

func GetBringFloatCommands(conn net.Conn) *cli.Command {
	return &cli.Command{
		Name:      "bring_float",
		Category:  "bring_float",
		ArgsUsage: "[mode]",
		Subcommands: []*cli.Command{
			{
				Name:        "current",
				Category:    "mode",
				Description: "Bring all lost floating windows of the current workspace to the center of the active monitor",
				Action: func(context *cli.Context) error {
					args := context.Args()

					if args.Len() > 0 {
						return fmt.Errorf("[ERROR] - this command do not take arguments")
					}

					if _, err := conn.Write([]byte("bring_float current")); err != nil {
						return fmt.Errorf("[ERROR] - %w\n", err)
					}

					buffer := make([]byte, 1024)
					_, err := conn.Read(buffer)
					if err != nil {
						return fmt.Errorf("[ERROR] - %w\n", err)
					}

					return nil
				},
			},
		},
	}
}
