package scratchpads

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
)

func GetScratchpadsCommands(conn net.Conn) *cli.Command {
	return &cli.Command{
		Name:      "scratchpads",
		Category:  "scratchpads",
		ArgsUsage: "[action]",
		Subcommands: []*cli.Command{
			{
				Name:      "toggle",
				Category:  "action",
				ArgsUsage: "[name]",
				Action: func(context *cli.Context) error {
					args := context.Args()

					if args.Len() > 1 {
						return fmt.Errorf("only 1 arg is necessary")
					}
					if _, err := conn.Write([]byte(fmt.Sprintf("scratchpads toggle %s", args.Get(0)))); err != nil {
						panic(err)
					}

					buffer := make([]byte, 1024)
					mLen, err := conn.Read(buffer)
					if err != nil {
						fmt.Println("Error reading:", err.Error())
						panic(err)
					}

					fmt.Println("Received: ", string(buffer[:mLen]))
					return nil
				},
			},
		},
	}
}
