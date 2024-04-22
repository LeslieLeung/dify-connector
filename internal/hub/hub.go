package hub

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/api"
	"github.com/leslieleung/dify-connector/internal/channel"
	"github.com/leslieleung/dify-connector/internal/command"
	"sync"
)

type Hub struct {
	channels []channel.Channel
	commands map[string]command.Command

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
	command.Commands = h.commands
	for _, c := range h.channels {
		h.wg.Add(1)
		go c.Start(ctx, h.wg)
	}
	h.wg.Add(1)
	go api.StartAPI(ctx, h.wg)
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
		if h.commands == nil {
			h.commands = make(map[string]command.Command)
		}
		for _, cmd := range commands {
			h.commands[cmd.GetName()] = cmd
		}
	}
}
