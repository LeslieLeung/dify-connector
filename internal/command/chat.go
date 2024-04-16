package command

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/leslieleung/dify-connector/internal/database"
	"github.com/leslieleung/dify-connector/pkg/dify"
)

type ChatCommand struct{}

var _ Command = (*ChatCommand)(nil)

func NewChatCommand() ChatCommand {
	return ChatCommand{}
}

func (c ChatCommand) GetName() string {
	return "chat"
}

func (c ChatCommand) GetDescription() string {
	return "chat with dify"
}

func (c ChatCommand) Execute(ctx context.Context, msg *Message) (string, error) {
	// get user state
	session, err := database.GetSession(ctx, msg.UserIdentifier)
	if err != nil {
		return "", err
	}
	// get dify config
	difyConf, err := database.GetDifyApp(ctx, session.State.CurrentApp)
	if err != nil {
		return "", err
	}
	if !difyConf.Enabled {
		return fmt.Sprintf("app %s is disabled", difyConf.Name), nil
	}
	difyApp := dify.New(difyConf.BaseURL, difyConf.APIKey)
	resp, err := difyApp.CompletionMessage(dify.CompletionMessageRequest{
		Inputs: map[string]interface{}{
			"query": msg.Body,
		},
		ResponseMode: dify.ResponseModeBlocking,
		User:         uuid.New().String(),
	})
	if err != nil {
		return "", err
	}
	return resp.Answer, nil
}
