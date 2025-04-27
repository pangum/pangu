package builder

import (
	"github.com/harluo/boot/internal/builder/internal/function"
	"github.com/harluo/boot/internal/internal/command"
	"github.com/harluo/boot/internal/internal/config"
)

type Command struct {
	params *config.Command
}

func NewCommand(name string) *Command {
	return &Command{
		params: config.NewCommand(name),
	}
}

func (c *Command) Build() *command.Default {
	return command.NewDefault(c.params)
}

func (c *Command) Aliases(required string, optionals ...string) *Command {
	return c.set(func() {
		c.params.Aliases = append(c.params.Aliases, required)
		c.params.Aliases = append(c.params.Aliases, optionals...)
	})
}

func (c *Command) Usage(usage string) *Command {
	return c.set(func() {
		c.params.Usage = usage
	})
}

func (c *Command) Category(category string) *Command {
	return c.set(func() {
		c.params.Category = category
	})
}

func (c *Command) Hidden() *Command {
	return c.set(func() {
		c.params.Hidden = true
	})
}

func (c *Command) set(set function.Set) (command *Command) {
	set()
	command = c

	return
}
