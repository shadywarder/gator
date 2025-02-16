package domain

// Commands struct holds all the commands CLI can handle.
type Commands struct {
	Callbacks map[string]func(*State, *Command) error
}

// NewCommands instantiates a new Commands entity.
func NewCommands() *Commands {
	return &Commands{
		Callbacks: make(map[string]func(*State, *Command) error),
	}
}

// Register registers a new command with specified callback.
func (c *Commands) Register(name string, f func(*State, *Command) error) {
	c.Callbacks[name] = f
}

// Call calls provided command.
func (c *Commands) Call(s *State, cmd *Command) error {
	f, exists := c.Callbacks[cmd.Name]
	if !exists {
		return ErrInvalidCommand
	}

	if err := f(s, cmd); err != nil {
		return err
	}

	return nil
}
