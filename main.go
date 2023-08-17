package main

import (
	"context"
	"github.com/edjubert/gophrland/cmd/gophrland"
)

func main() {
	ctx := context.Background()

	if err := gophrland.Run(ctx); err != nil {
		panic(err)
	}
}
