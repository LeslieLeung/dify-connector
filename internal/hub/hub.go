package hub

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/channel"
	"github.com/leslieleung/dify-connector/internal/command"
	"github.com/leslieleung/dify-connector/pkg/dify"
	"sync"
)

type Hub struct {
	channels []channel.Channel
	commands []command.Command
	difyApps []*dify.App

	wg *sync.WaitGroup
}

func New(opts ...Option) *Hub {
	h := &Hub{
		wg: &sync.WaitGroup{},
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

func (h *Hub) Start(ctx context.Context) {
	for _, c := range h.channels {
		println("Starting channel")
		h.wg.Add(1)
		go c.Start(ctx, h.wg)
	}
	h.wg.Wait()
}

type Option func(*Hub)

func RegisterChannels(channels ...channel.Channel) Option {
	return func(h *Hub) {
		h.channels = append(h.channels, channels...)
	}
}

func RegisterCommands(commands ...command.Command) Option {
	return func(h *Hub) {
		h.commands = append(h.commands, commands...)
	}
}

func RegisterDifyApps(apps ...*dify.App) Option {
	return func(h *Hub) {
		h.difyApps = append(h.difyApps, apps...)
	}
}
