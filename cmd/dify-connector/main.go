package difyconnector

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/channel"
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

	// TODO dynamically toggle channels
	channels := make([]channel.Channel, 0)
	if discordToken := os.Getenv("DISCORD_TOKEN"); discordToken != "" {
		channels = append(channels, channel.NewDiscord(discordToken))
	}
	if dingTalkClientID := os.Getenv("DINGTALK_CLIENT_ID"); dingTalkClientID != "" {
		channels = append(channels, channel.NewDingTalk(dingTalkClientID, os.Getenv("DINGTALK_CLIENT_SECRET")))
	}

	h := hub.New(
		hub.RegisterChannels(
			channels...,
		),
	)
	h.Start(ctx)
}
