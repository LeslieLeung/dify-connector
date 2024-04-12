package channel

import (
	"context"
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
