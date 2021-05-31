package command

import (
	`github.com/storezhang/pangu/app`
)

// Base 命令基类
type Base struct {
	// 名称
	name string
	// 别名
	aliases []string
	// 使用方法
	usage string
}

// NewBase 创建命令基类
func NewBase(name string, usage string, aliases ...string) Base {
	return Base{
		name:    name,
		aliases: aliases,
		usage:   usage,
	}
}

func (c *Base) Name() string {
	return c.name
}

func (c *Base) Aliases() []string {
	return c.aliases
}

func (c *Base) Usage() string {
	return c.usage
}

func (c *Base) SubCommands() (commands []app.Command) {
	return
}
