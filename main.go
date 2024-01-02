package main

import (
	"context"
	"makex/cmd"
)

func main() {
	ctx := context.Background()

	cmd.PrepareAndExecute(ctx)
}
