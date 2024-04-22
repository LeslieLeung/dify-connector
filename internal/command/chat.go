package command

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/leslieleung/dify-connector/internal/database"
	"github.com/leslieleung/dify-connector/internal/database/typedef"
	"github.com/leslieleung/dify-connector/pkg/dify"
)

type ChatCommand struct{}

var _ Command = (*ChatCommand)(nil)

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
	return generateResponse(difyConf, msg.Body), nil
}

func generateResponse(app *typedef.DifyApp, query string) string {
	difyApp := dify.New(app.BaseURL, app.APIKey)
	switch app.Type {
	case dify.AppTypeTextGenerator:
		resp, err := difyApp.CompletionMessage(dify.CompletionMessageRequest{
			Inputs: map[string]interface{}{
				"query": query,
			},
			ResponseMode: dify.ResponseModeBlocking,
			User:         uuid.New().String(),
		})
		if err != nil {
			return ""
		}
		return resp.Answer
	case dify.AppTypeChatApp:
		resp, err := difyApp.ChatMessageStream(dify.ChatMessageRequest{
			Inputs: map[string]interface{}{
				"query": query,
			},
			User: uuid.New().String(),
		})
		if err != nil {
			return ""
		}
		res, err := resp.Wait()
		if err != nil {
			return ""
		}
		return res
	default:
		return "Unknown app type"
	}
}
