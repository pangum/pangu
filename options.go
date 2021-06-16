package pangu

import (
	_ `embed`

	`github.com/storezhang/glog`
	`github.com/storezhang/pangu/app`
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
	// 作者
	authors []Author
	// 是否处理默认值
	isDefault bool
	// 是否验证数据
	isValidate bool
	// 日志
	logger app.Logger
}

//go:embed asset/template/help_app.tmpl
var helpAppTemplate string

//go:embed asset/template/help_command.tmpl
var helpCommandTemplate string

//go:embed asset/template/help_subcommand.tmpl
var helpSubcommandTemplate string

func defaultOptions() *options {
	return &options{
		usage: "一个功能强大的命令行应用程序框架",

		// 帮助信息
		helpAppTemplate:        helpAppTemplate,
		helpCommandTemplate:    helpCommandTemplate,
		helpSubcommandTemplate: helpSubcommandTemplate,

		// 版权
		copyright: "https://pangu.archtech.studio",
		authors: []Author{{
			Name:  "storezhang",
			Email: "storezhang@gmail.com",
		}, {
			Name:  "yangteng",
			Email: "yt443556@163.com",
		}},
		banner: banner{
			data:       "pangu",
			bannerType: BannerTypeAscii,
		},

		isDefault:  true,
		isValidate: true,
		logger:     glog.Must(glog.Zap()),
	}
}
