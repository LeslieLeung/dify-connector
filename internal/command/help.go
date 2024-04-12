package command

import "context"

type HelpCommand struct{}

var _ Command = (*HelpCommand)(nil)

func (c HelpCommand) GetName() string {
	return "help"
}

func (c HelpCommand) GetDescription() string {
	return "display help message"
}

func (c HelpCommand) Execute(_ context.Context, _ *Message) (string, error) {
	commands := getCommands()
	var help string
	for k, v := range commands {
		help += k + ": " + v + "\n"
	}
	return help, nil
}
