package cmd

import (
	"github.com/pangum/pangu/app"
)

// Command 命令
type Command struct {
	// 名称
	name string
	// 别名
	aliases []string
	// 使用方法
	usage string
}

// New 创建命令
func New(name string, opts ...option) Command {
	_options := defaultOption()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return Command{
		name:    name,
		aliases: _options.aliases,
		usage:   _options.usage,
	}
}

func (c *Command) Name() string {
	return c.name
}

func (c *Command) Aliases() []string {
	return c.aliases
}

func (c *Command) Usage() string {
	return c.usage
}

func (c *Command) Args() (args []app.Arg) {
	return
}

func (c *Command) SubCommands() (commands []app.Command) {
	return
}
