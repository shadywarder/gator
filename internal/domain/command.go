package domain

// Command stores the command's necessary data.
type Command struct {
	Name string
	Args []string
}

// NewCommand instantiates a new Command entity.
func NewCommand(name string, args []string) *Command {
	return &Command{
		Name: name,
		Args: args,
	}
}
