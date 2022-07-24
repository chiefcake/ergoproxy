package main

import (
	"context"

	"github.com/chiefcake/ergoproxy/internal/app"
)

// @title ErgoProxy API
// @version 1.0

// @host localhost:8080
func main() {
	ctx := context.Background()

	app.Run(ctx)
}
