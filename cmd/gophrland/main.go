// socket-server project main.go
package main

import (
	"context"
	"github.com/edjubert/gophrland/cmd/gophrland/cmd"
	"github.com/edjubert/gophrland/pkg/logging"
	"os"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	_ = logging.New(os.Stdout)

	//ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	//defer cancel()

	return cmd.Execute()
}
