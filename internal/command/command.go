package command

import "context"

type Command interface {
	GetName() string
	GetDescription() string
	// Execute the command with the given argument and generate a response(optional)
	Execute(ctx context.Context, msg *Message) (string, error)
}

type Message struct {
	Command        string
	Body           string
	UserIdentifier string // platform
}

var Commands map[string]Command

func Process(ctx context.Context, msg *Message) (string, error) {
	cmd, ok := Commands[msg.Command]
	if !ok {
		return "", nil
	}
	return cmd.Execute(ctx, msg)
}

func getCommands() map[string]string {
	cmds := make(map[string]string)
	for k, v := range Commands {
		cmds[k] = v.GetDescription()
	}
	return cmds
}

func IsCommand(input string) bool {
	_, ok := Commands[input]
	return ok
}
