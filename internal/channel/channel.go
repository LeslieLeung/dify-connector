package channel

import (
	"context"
	"fmt"
	"github.com/leslieleung/dify-connector/internal/database"
	"github.com/leslieleung/dify-connector/internal/database/typedef"
	"sync"
)

type Channel interface {
	// Start the channel
	Start(ctx context.Context, wg *sync.WaitGroup)
	// Stop the channel and clean up
	Stop(ctx context.Context)
}

const (
	TypeDiscord = iota
	TypeDingTalk
)

func LoadChannels(ctx context.Context) ([]Channel, error) {
	channels, err := database.GetEnabledChannels(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]Channel, 0)
	for _, channel := range channels {
		c, err := buildChannel(channel)
		if err != nil {
			return nil, err
		}
		out = append(out, c)
	}

	return out, nil
}

func buildChannel(channel *typedef.Channel) (Channel, error) {
	var (
		c   Channel
		err error
	)
	switch channel.Type {
	case TypeDiscord:
		c, err = NewDiscordWithCredential(channel.Credential)
	case TypeDingTalk:
		c, err = NewDingTalkWithCredential(channel.Credential)
	default:
		return nil, fmt.Errorf("unsupported channel type: %d", channel.Type)
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}
