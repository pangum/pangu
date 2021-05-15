package pangu

import (
	_ `embed`
)

type options struct {
	// 应用名称
	name string
	// 应用描述
	description string
	// 使用方式
	usage string
	// 帮助信息模板
	helpAppTemplate string
	// 命令帮助信息模板
	helpCommandTemplate string
	// 子命令帮助信息模板
	helpSubcommandTemplate string
	// 标志
	banner banner
	// 版权
	copyright string
	// 是否处理默认值
	isDefault bool
	// 是否验证数据
	isValidate bool
}

//go:embed asset/template/help_app.tmpl
var helpAppTemplate string

//go:embed asset/template/help_command.tmpl
var helpCommandTemplate string

//go:embed asset/template/help_subcommand.tmpl
var helpSubcommandTemplate string

func defaultOptions() options {
	return options{
		usage:                  "一个新式的命令行程序",
		helpAppTemplate:        helpAppTemplate,
		helpCommandTemplate:    helpCommandTemplate,
		helpSubcommandTemplate: helpSubcommandTemplate,

		isDefault:  true,
		isValidate: true,
	}
}
