package server

import (
	"fmt"
	"gophrland/cmd/internal/client"
	"gopkg.in/yaml.v3"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

type Scratchpad struct {
	Command   string `yaml:"command"`
	Animation string `yaml:"animation"`
	Unfocus   string `yaml:"unfocus"`
}

type Options struct {
	Scratchpads []map[string]Scratchpad `yaml:"scratchpads"`
}

type Config struct {
	Plugins []string `yaml:"plugins"`
	Options Options  `yaml:"options"`
}

const DEFAULT_CONFIG = "./config.yaml"

func ReadConfig(file string) {
	dat, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Could not read file '%s' -> %v\n", file, err)
	}

	var config Config
	if err := yaml.Unmarshal(dat, &config); err != nil {
		fmt.Printf("Could not unmarshal %v\n", err)
	}

	fmt.Println(config.Plugins)
	fmt.Println(config.Options)
}
func closeServer(server net.Listener) {
	if err := server.Close(); err != nil {
		panic(err)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("Received: ", string(buffer[:mLen]))
	if _, err := connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen]))); err != nil {
		panic(err)
	}

	if err := connection.Close(); err != nil {
		panic(err)
	}
}

func Handle() {
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		client.Handle()
		os.Exit(1)
	}
	defer closeServer(server)

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		fmt.Println("client connected")
		go processClient(connection)
	}
}
