package server

import (
	"context"
	"github.com/edjubert/gophrland/pkg/logging"
	serverCmd "github.com/edjubert/gophrland/pkg/server/cmd"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

type appArgs struct {
	Port   int
	Config string
}

func AddCommand(cmd *cobra.Command, config string) {
	var args appArgs
	daemon := &cobra.Command{
		Use:   "daemon",
		Short: "Run the Gophrland server",
		Long:  "The Gophrland local server instanciate the gophrland session",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, pargs []string) error {
			return run(cmd.Context(), args)
		},
	}

	fl := daemon.PersistentFlags()
	fl.SortFlags = false

	fl.StringVarP(
		&args.Config, "config", "c", config,
		"config file (default ~/.config/hyprland/gophrland.yaml)",
	)

	fl.IntVarP(
		&args.Port, "port", "p", 9988,
		"The port to use for local TCP server",
	)

	cmd.AddCommand(daemon)
}

func run(ctx context.Context, args appArgs) error {
	logger := logging.New(os.Stdout)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return serverCmd.New(
			serverCmd.WithLogger(logger),
			serverCmd.WithConfigFilePath(args.Config),
			serverCmd.WithPort(args.Port),
		)
	})

	return eg.Wait()
}
