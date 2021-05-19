package app

// Command 描述一个可以被执行的命令
type Command interface {
	parameter

	// Run 命令实际执行
	Run(ctx *Context) error

	// SubCommands 子命令
	SubCommands() []Command
}
