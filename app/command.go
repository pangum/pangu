package app

// Command 描述一个可以被执行的命令
type Command interface {
	parameter

	// Run 执行命令
	Run(ctx *Context) error

	// Args 参数列表
	Args() []Arg

	// Subcommands 子命令列表
	Subcommands() []Command

	// Description 描述
	Description() string
}
