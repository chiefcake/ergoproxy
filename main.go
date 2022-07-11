package main

import (
	"context"

	"github.com/chiefcake/ergoproxy/internal/gateway"
)

func main() {
	ctx := context.Background()

	gateway.Run(ctx)
}
