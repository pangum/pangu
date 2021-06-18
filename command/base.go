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

func (b *Base) Name() string {
	return b.name
}

func (b *Base) Aliases() []string {
	return b.aliases
}

func (b *Base) Usage() string {
	return b.usage
}

func (b *Base) SubCommands() (commands []app.Command) {
	return
}
