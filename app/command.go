package app

// Command 描述一个可以被执行的命令
type Command interface {
	parameter

	// Run 执行命令
	Run(ctx *Context) error

	// Arguments 参数列表
	Arguments() Arguments

	// Subcommands 子命令列表
	Subcommands() Commands

	// Description 描述文字
	Description() string

	// Category 分类
	Category() string
}
