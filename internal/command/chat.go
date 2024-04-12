package command

import (
	"github.com/google/uuid"
	"github.com/leslieleung/dify-connector/pkg/dify"
	"os"
)

type ChatCommand struct {
}

func NewChatCommand() ChatCommand {
	return ChatCommand{}
}

func (c ChatCommand) GetName() string {
	return "chat"
}

func (c ChatCommand) Execute(arg string) (string, error) {
	// TODO better way to handle dify apps
	difyApp := dify.New(os.Getenv("DIFY_BASE_URL"), os.Getenv("DIFY_API_KEY"))
	resp, err := difyApp.CompletionMessage(dify.CompletionMessageRequest{
		Inputs: map[string]interface{}{
			"query": arg,
		},
		ResponseMode: dify.OutputModeBlocking,
		User:         uuid.New().String(),
	})
	if err != nil {
		return "", err
	}
	return resp.Answer, nil
}
