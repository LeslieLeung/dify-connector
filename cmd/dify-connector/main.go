package difyconnector

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/channel"
	"github.com/leslieleung/dify-connector/internal/command"
	"github.com/leslieleung/dify-connector/internal/database"
	"github.com/leslieleung/dify-connector/internal/hub"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Run: runServe,
}

func runServe(_ *cobra.Command, _ []string) {
	ctx, cancel := context.WithCancel(context.Background())

	// handle exit signals gracefully
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		if <-sigChan; true {
			slog.Info("Shutting down")
			cancel()
		}
	}()

	if os.Getenv("DATABASE_DSN") == "" {
		println("DATABASE_DSN is required")
		os.Exit(1)
	}

	// Initialize DB
	database.Init(os.Getenv("DATABASE_DSN"))

	channels, err := channel.LoadChannels(ctx)
	if err != nil {
		panic(err)
	}

	// TODO just a temporary thing, need to find a better way to initialize
	if len(channels) == 0 {
		slog.Info("No channels to start")
		return
	}

	h := hub.New(
		hub.RegisterChannels(
			channels...,
		),
		hub.RegisterCommands(
			command.ChatCommand{},
			command.AppCommand{},
			command.HelpCommand{},
		),
	)
	h.Start(ctx)
}
