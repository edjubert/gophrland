package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/pkg/logging"
	internal2 "github.com/edjubert/gophrland/pkg/server/internal"
	"github.com/edjubert/gophrland/pkg/server/pkg/IPC"
	"github.com/edjubert/gophrland/pkg/server/pkg/config"
	"github.com/edjubert/gophrland/plugins"
	"os"
)

type serverOptions struct {
	logger     logging.Logger
	port       int
	host       string
	configPath string
}

type Option func(opts *serverOptions)

func WithLogger(logger logging.Logger) Option {
	return func(opts *serverOptions) {
		opts.logger = logger
	}
}

func WithPort(port int) Option {
	return func(opts *serverOptions) {
		opts.port = port
	}
}

func WithConfigFilePath(path string) Option {
	return func(opts *serverOptions) {
		opts.configPath = path
	}
}

const HyprlandInstanceSignature = "HYPRLAND_INSTANCE_SIGNATURE"

func New(options ...Option) error {
	opts := serverOptions{logger: logging.Noop, port: 9988, host: plugins.ServerHost}

	for _, opt := range options {
		opt(&opts)
	}

	server := internal2.CreateServer(opts.host, opts.port)
	defer internal2.CloseServer(server)

	loadedConf := config.ReadConfig(opts.configPath)
	plugins.ApplyConfig(loadedConf)

	hyprlandSignature := os.Getenv(HyprlandInstanceSignature)
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

	return nil
}
