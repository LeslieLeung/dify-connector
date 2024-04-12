package message

import "github.com/leslieleung/dify-connector/internal/command"

type Message struct {
	Command string
	Body    string
}

var commands = map[string]command.Command{
	"chat": command.NewChatCommand(),
}

func Process(msg *Message) (string, error) {
	cmd, ok := commands[msg.Command]
	if !ok {
		return "", nil
	}
	return cmd.Execute(msg.Body)
}
