package bootstrap

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/leslieleung/dify-connector/internal/channel"
	"github.com/leslieleung/dify-connector/internal/database"
	"github.com/leslieleung/dify-connector/internal/database/typedef"
	"strings"
)

// BuildChannel creates a channel based on the bootstrap string
// boostrap the first channel
// <type>:<credential1>:<credential2>:...
// dingtalk:client_id:client_secret
// discord:token
func BuildChannel(ctx context.Context, bootstrap string) (channel.Channel, error) {
	var (
		c          channel.Channel
		credential any
	)
	parts := strings.Split(bootstrap, ":")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid bootstrap string: %s", bootstrap)
	}

	switch parts[0] {
	case channel.TypeStrDingTalk:
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid dingtalk bootstrap string: %s", bootstrap)
		}
		c = channel.NewDingTalk(parts[1], parts[2])
		credential = &channel.DingTalkCredential{
			ClientID:     parts[1],
			ClientSecret: parts[2],
		}
	case channel.TypeStrDiscord:
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid discord bootstrap string: %s", bootstrap)
		}
		c = channel.NewDiscord(parts[1])
		credential = &channel.DiscordCredential{
			Token: parts[1],
		}
	default:
		return nil, fmt.Errorf("unsupported channel type: %s", parts[0])
	}

	// serialize the credential
	credentialBytes, err := json.Marshal(credential)
	if err != nil {
		return nil, err
	}

	// store the channel
	err = database.SaveChannel(ctx, &typedef.Channel{
		Name:       "Default",
		Type:       channel.TypeMap[parts[0]],
		Credential: string(credentialBytes),
		Enabled:    false,
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}
