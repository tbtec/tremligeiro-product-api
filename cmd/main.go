package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver/server"
)

func main() {

	ctx := context.Background()

	if err := run(ctx); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run(ctx context.Context) error {

	config, err := env.LoadEnvConfig()
	if err != nil {
		log.Fatal(err)
	}

	container, err := container.New(config)
	if err != nil {
		log.Fatal(err)
	}

	errStart := container.Start()
	if errStart != nil {
		log.Fatal(err)
	}

	httpServer := server.New(container, config)
	httpServer.Listen()

	slog.InfoContext(ctx, "Shutting down services...")

	return nil
}
