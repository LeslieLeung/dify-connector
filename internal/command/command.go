package command

import "fmt"

type Command interface {
	GetName() string
	Execute(arg string) (string, error)
}

var commandsMap = make(map[string]Command)

func RegisterCommands(commands ...Command) {
	for _, command := range commands {
		commandsMap[command.GetName()] = command
	}
}

func ExecuteCommand(name, arg string) (string, error) {
	command, ok := commandsMap[name]
	if !ok {
		return "", fmt.Errorf("command not found: %s", name)
	}
	return command.Execute(arg)
}
