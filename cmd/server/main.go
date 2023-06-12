package server

import (
	"fmt"
	"gophrland/cmd/server/cmd/IPC"
	"os"

	"gophrland/cmd/server/cmd/config"
	"gophrland/cmd/server/internal"
)

const DEFAULT_CONFIG = "./config.yaml"
const HYPRLAND_INSTANCE_SIGNATURE = "HYPRLAND_INSTANCE_SIGNATURE"

func Handle() {

	server := internal.CreateServer()
	defer internal.CloseServer(server)

	loadedConf := config.ReadConfig(DEFAULT_CONFIG)
	fmt.Println("loaded conf: ", loadedConf)
	config.ApplyConfig(loadedConf)

	hyprlandSignature := os.Getenv(HYPRLAND_INSTANCE_SIGNATURE)
	fmt.Println("signature: ", hyprlandSignature)
	go IPC.ConnectEvents(hyprlandSignature)
	//_, err := IPC.ConnectHyprctl(hyprlandSignature)
	//if err != nil {
	//	fmt.Println("[ERROR] - Could not open hyprctl socket", err)
	//}

	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go internal.ProcessClient(connection, loadedConf)
	}
}
