package cmd

import (
	"github.com/goexl/gox"
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
	// 描述
	description string

	_ gox.CannotCopy
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

func (c *Command) Run(_ *app.Context) (err error) {
	return
}

func (c *Command) Arguments() (args app.Arguments) {
	return
}

func (c *Command) Subcommands() (commands app.Commands) {
	return
}

func (c *Command) Description() string {
	return c.description
}
