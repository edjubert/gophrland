package cmd

import (
	"fmt"
	"github.com/edjubert/gophrland/pkg/client/pkg/tools"
	"github.com/edjubert/gophrland/pkg/logging"
	server "github.com/edjubert/gophrland/pkg/server/internal"
	"github.com/edjubert/gophrland/pkg/server/pkg/IPC"
	"github.com/edjubert/gophrland/pkg/server/pkg/config"
	"github.com/edjubert/gophrland/plugins"
	"os"
)

type serverOptions struct {
	logger     logging.Logger
	configPath string
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

func New(options ...Option) error {
	opts := serverOptions{logger: logging.Noop}

	for _, opt := range options {
		opt(&opts)
	}

	hyprlandSignature := tools.GetSignature()
	s := server.CreateSocket(hyprlandSignature)
	defer func() {
		_ = s.Close()
	}()

	loadedConf := config.ReadConfig(opts.configPath)
	plugins.ApplyConfig(loadedConf)

	go IPC.ConnectEvents(hyprlandSignature)
	//_, err := IPC.ConnectHyprctl(hyprlandSignature)
	//if err != nil {
	//	fmt.Println("[ERROR] - Could not open hyprctl socket", err)
	//}

	fmt.Println("[INFO] - Waiting for client...")
	for {
		connection, err := s.Accept()
		if err != nil {
			fmt.Println("[ERROR] - Error accepting: ", err.Error())
			os.Exit(1)
		}

		go server.ProcessClient(connection, loadedConf)
	}

	return nil
}
