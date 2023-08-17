// socket-server project main.go
package gophrland

import (
	"context"
	"github.com/edjubert/gophrland/cmd/gophrland/cmd"
	"github.com/edjubert/gophrland/pkg/logging"
	"os"
)

func Run(ctx context.Context) error {
	_ = logging.New(os.Stdout)

	return cmd.Execute()
}
