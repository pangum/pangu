package app

import (
	"github.com/pangum/pangu/internal/runtime"
	"github.com/urfave/cli/v2"
)

type Commands []Command

func (c Commands) Cli() (commands []*cli.Command) {
	commands = make([]*cli.Command, 0, len(c))
	for _, command := range c {
		cloned := command
		commands = append(commands, &cli.Command{
			Name:        cloned.Name(),
			Aliases:     cloned.Aliases(),
			Usage:       cloned.Usage(),
			Description: cloned.Description(),
			Subcommands: cloned.Subcommands().Cli(),
			Category:    cloned.Category(),
			Flags:       cloned.Arguments().Flags(),
			Hidden:      cloned.Hidden(),
			Action:      c.action(cloned),
		})
	}

	return
}

func (c Commands) action(command Command) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		return command.Run(runtime.NewContext(ctx))
	}
}
