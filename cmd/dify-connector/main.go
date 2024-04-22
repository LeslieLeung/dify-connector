package difyconnector

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/bootstrap"
	"github.com/leslieleung/dify-connector/internal/channel"
	"github.com/leslieleung/dify-connector/internal/command"
	"github.com/leslieleung/dify-connector/internal/database"
	"github.com/leslieleung/dify-connector/internal/hub"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var ServeCmd = &cobra.Command{
	Use: "serve",
	Run: runServe,
}

func runServe(_ *cobra.Command, _ []string) {
	if os.Getenv("DATABASE_DSN") == "" {
		println("DATABASE_DSN is required")
		os.Exit(1)
	}

	ctx := context.Background()
	// handle exit signals gracefully
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initialize DB
	database.Init(os.Getenv("DATABASE_DSN"))

	channels, err := channel.LoadChannels(ctx)
	if err != nil {
		panic(err)
	}

	if len(channels) == 0 {
		// read from BOOTSTRAP_CHANNEL
		bs := os.Getenv("BOOTSTRAP_CHANNEL")
		if bs == "" {
			println("No channels found and BOOTSTRAP_CHANNEL is not set")
			os.Exit(1)
		}
		// build channel
		c, err := bootstrap.BuildChannel(ctx, bs)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}
		channels = append(channels, c)
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
