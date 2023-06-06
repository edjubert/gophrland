package server

import (
	"fmt"
	"os"

	"gophrland/cmd/server/cmd/config"
	"gophrland/cmd/server/internal"
)

const DEFAULT_CONFIG = "./config.yaml"

func Handle() {

	server := internal.CreateServer()
	defer internal.CloseServer(server)

	loadedConf := config.ReadConfig(DEFAULT_CONFIG)
	config.ApplyConfig(loadedConf)

	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go internal.ProcessClient(connection)
	}
}
