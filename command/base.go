package command

import (
	`github.com/storezhang/pangu/app`
)

// Base 命令基类
type Base struct {
	// 名称
	Name string
	// 别名
	Aliases []string
	// 使用方法
	Usage string
}

func (c *Base) GetName() string {
	return c.Name
}

func (c *Base) GetAliases() []string {
	return c.Aliases
}

func (c *Base) GetUsage() string {
	return c.Usage
}

func (c *Base) SubCommands() (commands []app.Command) {
	return
}
