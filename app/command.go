package app

import (
	`github.com/storezhang/pangu/command`
)

// Command 描述一个可以被执行的命令
type Command interface {
	flag

	// Run 命令实际执行
	Run(ctx *command.Context) error

	// SubCommands 子命令
	SubCommands() []Command
}
