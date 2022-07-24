package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chiefcake/ergoproxy/internal/config"
	"github.com/chiefcake/ergoproxy/internal/handler"
	"github.com/chiefcake/ergoproxy/internal/model"
	"github.com/chiefcake/ergoproxy/internal/server"
	"github.com/chiefcake/ergoproxy/internal/service"
	"github.com/chiefcake/ergoproxy/internal/storage"
)

const timeout = 3 * time.Second

// Run starts the app gateway with provided config values.
func RunWithConfig(ctx context.Context, cfg *config.Config) {
	mapStorage := storage.NewMap[*model.RedirectRequest, model.RedirectResponse]()
	proxyService := service.NewProxy(mapStorage)
	proxyHandler := handler.NewProxy(proxyService)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	server := server.New(ctx, cfg, proxyHandler)

	log.Printf("Server is running on [%s:%s]...\n", cfg.ServerHost, cfg.ServerPort)

	go func() {
		err := server.Serve()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	<-shutdown

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	log.Println("Shutting down server...")

	err := server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
		return
	}
}

// Run parses the config values and starts the app gateway.
func Run(ctx context.Context) {
	config, err := config.New()
	if err != nil {
		log.Println(err)
		return
	}

	RunWithConfig(ctx, config)
}
