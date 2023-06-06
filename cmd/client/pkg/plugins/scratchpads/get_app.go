package scratchpads

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"net"
)

func GetApp(conn net.Conn) *cli.Command {
	return &cli.Command{
		Name:      "doo",
		Aliases:   []string{"do"},
		Category:  "motion",
		ArgsUsage: "[arrgh]",
		Subcommands: []*cli.Command{
			{
				Name: "wop",
				Action: func(context *cli.Context) error {
					fmt.Println("wop")

					if _, err := conn.Write([]byte("Hello Server! Greetings.")); err != nil {
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
