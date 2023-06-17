package cmd

import (
	"fmt"
	internal2 "gophrland/pkg/server/internal"
	"gophrland/pkg/server/pkg/IPC"
	"gophrland/pkg/server/pkg/config"
	"os"
)

const DEFAULT_CONFIG = "/home/edouard.jubert.ext/.config/hypr/gophrland.yaml"
const HYPRLAND_INSTANCE_SIGNATURE = "HYPRLAND_INSTANCE_SIGNATURE"

func New() {

	server := internal2.CreateServer()
	defer internal2.CloseServer(server)

	loadedConf := config.ReadConfig(DEFAULT_CONFIG)
	config.ApplyConfig(loadedConf)

	hyprlandSignature := os.Getenv(HYPRLAND_INSTANCE_SIGNATURE)
	go IPC.ConnectEvents(hyprlandSignature)
	//_, err := IPC.ConnectHyprctl(hyprlandSignature)
	//if err != nil {
	//	fmt.Println("[ERROR] - Could not open hyprctl socket", err)
	//}

	fmt.Println("[INFO] - Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("[ERROR] - Error accepting: ", err.Error())
			os.Exit(1)
		}

		go internal2.ProcessClient(connection, loadedConf)
	}
}
