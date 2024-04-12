package command

type HelpCommand struct {
}

func (c *HelpCommand) GetName() string {
	return "help"
}

func (c *HelpCommand) Execute(arg string) (string, error) {
	return "help", nil
}
