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
			fmt.Println("Reading expose")
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
					fmt.Println("Reading expose toggle")
					args := context.Args()

					if args.Len() > 0 {
						return fmt.Errorf("only 1 arg is necessary")
					}
					if _, err := conn.Write([]byte(fmt.Sprintf("expose toggle %s", args.Get(0)))); err != nil {
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
