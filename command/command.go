package command

import (
	`github.com/storezhang/pangu/app`
)

// Command 命令基类
type Command struct {
	// 名称
	Name string
	// 别名
	Aliases []string
	// 使用方法
	Usage string
}

func (c *Command) GetName() string {
	return c.Name
}

func (c *Command) GetAliases() []string {
	return c.Aliases
}

func (c *Command) GetUsage() string {
	return c.Usage
}

func (c *Command) SubCommands() (commands []app.Command) {
	return
}
