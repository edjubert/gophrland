package expose

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
)

func GetExposeCommands(conn net.Conn) *cli.Command {
	return &cli.Command{
		Name:      "expose",
		Category:  "expose",
		ArgsUsage: "[action]",
		Action: func(context *cli.Context) error {
			if _, err := conn.Write([]byte("expose")); err != nil {
				fmt.Println("[ERROR] - ", err)
				panic(err)
			}
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:     "toggle",
				Category: "action",
				Action: func(context *cli.Context) error {
					args := context.Args()

					if args.Len() > 0 {
						return fmt.Errorf("[ERROR] -  only 1 arg is necessary")
					}
					if _, err := conn.Write([]byte(fmt.Sprintf("expose toggle %s", args.Get(0)))); err != nil {
						panic(err)
					}

					buffer := make([]byte, 1024)
					_, err := conn.Read(buffer)
					if err != nil {
						fmt.Println("[ERROR] - Error reading:", err.Error())
						panic(err)
					}

					return nil
				},
			},
		},
	}
}
