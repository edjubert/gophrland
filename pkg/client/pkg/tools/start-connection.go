package tools

import (
	"net"
	"os"
)

const HyprlandInstanceSignature = "HYPRLAND_INSTANCE_SIGNATURE"

func GetSignature() string {
	return os.Getenv(HyprlandInstanceSignature)
}

func StartUnixConnection() net.Conn {
	connection, err := net.Dial("unix", "/tmp/hypr/"+GetSignature()+"/.gophrland.sock")
	if err != nil {
		panic(err)
	}

	return connection
}
