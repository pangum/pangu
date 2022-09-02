package pangu

import (
	_ "embed"

	"github.com/goexl/simaqian"
	"github.com/pangum/pangu/app"
)

//go:embed asset/template/help_app.tmpl
var helpAppTemplate string

//go:embed asset/template/help_command.tmpl
var helpCommandTemplate string

//go:embed asset/template/help_subcommand.tmpl
var helpSubcommandTemplate string

type (
	option interface {
		apply(options *options)
	}

	options struct {
		*configOptions

		// 合法性验证，包括
		// 启动器构造方法合法性验证
		// 构造方法合法性验证
		verify bool
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
		banner _banner
		// 版权
		copyright string
		// 作者
		authors []*author

		// 日志
		logger app.Logger
	}
)

func defaultOptions() *options {
	return &options{
		configOptions: defaultConfigOptions(),

		verify:      true,
		description: `一个基于Pangum/Pangu快速开始框架的应用程序`,
		usage:       `一个功能强大的命令行应用程序框架`,

		// 帮助信息
		helpAppTemplate:        helpAppTemplate,
		helpCommandTemplate:    helpCommandTemplate,
		helpSubcommandTemplate: helpSubcommandTemplate,

		// 版权
		copyright: copyright,
		authors: []*author{{
			name:  authorName,
			email: authorEmail,
		}},
		banner: _banner{
			data: `pangu`,
			typ:  BannerTypeAscii,
		},

		logger: simaqian.Must(simaqian.Zap()),
	}
}
