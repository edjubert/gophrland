package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/pkg/logging"
	server "github.com/edjubert/gophrland/pkg/server/internal"
	"github.com/edjubert/gophrland/pkg/server/pkg/config"
	"github.com/edjubert/gophrland/plugins"
	IPC "github.com/edjubert/hyprland-ipc-go/ipc"
	"os"
)

type serverOptions struct {
	logger     logging.Logger
	configPath string
	restart    bool
}

type Option func(opts *serverOptions)

func WithLogger(logger logging.Logger) Option {
	return func(opts *serverOptions) {
		opts.logger = logger
	}
}

func WithConfigFilePath(path string) Option {
	return func(opts *serverOptions) {
		opts.configPath = path
	}
}

func WithRestart(restart bool) Option {
	return func(opts *serverOptions) {
		opts.restart = restart
	}
}

const UnixSocketName = ".gophrland.sock"

func New(options ...Option) error {
	opts := serverOptions{logger: logging.Noop}

	for _, opt := range options {
		opt(&opts)
	}

	if opts.restart {
		if err := IPC.RemoveSocket(UnixSocketName); err != nil {
			fmt.Printf("[ERROR] - Could not remove UnixSocketName %s\n", UnixSocketName)
		}
	}

	s := IPC.CreateSocket(UnixSocketName)
	defer func() {
		_ = s.Close()
	}()

	loadedConf := config.ReadConfig(opts.configPath)
	plugins.ApplyConfig(loadedConf)

	fmt.Println("[INFO] - Waiting for client...")
	for {
		connection, err := s.Accept()
		if err != nil {
			fmt.Println("[ERROR] - Error accepting: ", err.Error())
			os.Exit(1)
		}

		go server.ProcessClient(connection, loadedConf)
	}
}
