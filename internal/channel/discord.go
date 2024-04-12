package channel

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/leslieleung/dify-connector/internal/message"
	"log/slog"
	"sync"
)

type Discord struct {
	Token string

	dg *discordgo.Session
}

func NewDiscord(token string) *Discord {
	return &Discord{
		Token: token,
	}
}

func (d *Discord) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	// Create a new Discord session using the provided bot token.
	var err error
	d.dg, err = discordgo.New("Bot " + d.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	fmt.Println("Discord session created")

	// Register the messageCreate func as a callback for MessageCreate events.
	d.dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	d.dg.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

	// Open a websocket connection to Discord and begin listening.
	err = d.dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Run a loop until the context is done
	for {
		if <-ctx.Done(); true {
			// If the context is done, Stop the Discord session and return
			d.Stop(ctx)
			println("Discord session closed")
			return
		}
	}
}

func (d *Discord) Stop(_ context.Context) {
	d.dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	slog.Info("messageCreate", "msg", m.Content)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// if the bot is not in the mention list, ignore the message
	if !s.State.User.Bot {
		return
	}

	// send the message through dify app and get response
	msg := &message.Message{
		Command: "chat",
		Body:    m.Content,
	}
	resp, err := message.Process(msg)
	if err != nil {
		slog.Error("failed to process message", "err", err)
		return
	}
	s.ChannelMessageSend(m.ChannelID, resp)
}
