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
	helpTemplate string
	// 标志
	banner banner
	// 是否处理默认值
	isDefault bool
	// 是否验证数据
	isValidate bool
}

//go:embed asset/template/help.txt
var helpTemplate string

func defaultOptions() options {
	return options{
		usage:        "一个新式的命令行程序",
		helpTemplate: helpTemplate,

		isDefault:  true,
		isValidate: true,
	}
}
