package channel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/leslieleung/dify-connector/internal/command"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/chatbot"
	"github.com/open-dingtalk/dingtalk-stream-sdk-go/client"
	"log/slog"
	"strings"
	"sync"
)

type DingTalk struct {
	ClientID     string
	ClientSecret string

	dt *client.StreamClient
}

var _ Channel = (*DingTalk)(nil)

type DingTalkCredential struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func NewDingTalk(clientID, clientSecret string) *DingTalk {
	return &DingTalk{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func NewDingTalkWithCredential(credential string) (*DingTalk, error) {
	cred := &DingTalkCredential{}
	err := json.Unmarshal([]byte(credential), cred)
	if err != nil {
		return nil, err
	}
	return NewDingTalk(cred.ClientID, cred.ClientSecret), nil
}

func (d *DingTalk) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	var err error
	d.dt = client.NewStreamClient(client.WithAppCredential(client.NewAppCredentialConfig(d.ClientID, d.ClientSecret)))
	d.dt.RegisterChatBotCallbackRouter(onChatBotMessageReceived)

	err = d.dt.Start(ctx)
	if err != nil {
		fmt.Println("error starting DingTalk stream client,", err)
		return
	}
	fmt.Println("DingTalk stream client started")

	// Run a loop until the context is done
	for {
		if <-ctx.Done(); true {
			// If the context is done, Stop the Discord session and return
			d.Stop(ctx)
			println("DingTalk stream client closed")
			return
		}
	}
}

func (d *DingTalk) Stop(_ context.Context) {
	d.dt.Close()
}

const DingTalkBotPrefix = DiscordBotPrefix

func onChatBotMessageReceived(ctx context.Context, data *chatbot.BotCallbackDataModel) ([]byte, error) {
	slog.Info("Received message from DingTalk chatbot", "sender", data.SenderId, "content", data.Text.Content, "chatbotUserId", data.ChatbotUserId)
	// if the message is from the bot itself, ignore it
	if data.SenderId == data.ChatbotUserId {
		return []byte(""), nil
	}
	// TODO if the bot is not on the mentioned list, ignore it

	content := strings.TrimSpace(strings.TrimPrefix(data.Text.Content, fmt.Sprintf(DingTalkBotPrefix, data.ChatbotUserId)))
	parts := strings.Fields(content)

	if len(parts) == 0 {
		return []byte(""), nil
	}

	msg := &command.Message{
		Command:        "chat",
		Body:           content,
		UserIdentifier: "dingtalk:" + data.SenderId,
	}
	if command.IsCommand(parts[0]) {
		msg.Command = strings.TrimSpace(parts[0])
		msg.Body = strings.TrimSpace(strings.TrimPrefix(content, parts[0]))
	}
	resp, err := command.Process(ctx, msg)
	if err != nil {
		slog.Error("Error processing message", "error", err)
		return nil, err
	}

	replier := chatbot.NewChatbotReplier()
	if err := replier.SimpleReplyText(ctx, data.SessionWebhook, []byte(resp)); err != nil {
		return nil, err
	}
	return []byte(""), nil
}
