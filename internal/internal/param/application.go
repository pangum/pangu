package param

import (
	"github.com/heluon/boot/internal/internal/constant"
	"github.com/heluon/boot/internal/internal/kernel"
	"github.com/heluon/boot/internal/internal/loader"
)

type Application struct {
	// 配置
	Config *Config
	// 帮助
	Help *Help
	// 徽标
	Banner *Banner
	// 状态码
	Code *Code
	// 超时
	Timeout *Timeout

	// 合法性验证，包括
	// 启动器构造方法合法性验证
	// 构造方法合法性验证
	Validate bool
	// 应用描述
	Description string
	// 使用方式
	Usage string

	// 版权
	Copyright string
	// 作者
	Authors kernel.Authors
	// 元数据
	Metadata map[string]any
}

func NewApplication() *Application {
	return &Application{
		Config: NewConfig(
			loader.NewJson(),
			loader.NewXml(),
		),
		Help:    newHelp(),
		Banner:  newBanner(),
		Code:    newCode(),
		Timeout: NewTimeout(),

		Validate:    true,
		Description: "一个使用github.com/pangum/pangu构建的应用程序，可以使用应用程序提供的命令来使用本程序",
		Usage:       "一个使用github.com/pangum/pangu构建的应用程序",

		Copyright: constant.Copyright,
		Authors: kernel.Authors{{
			Name:  constant.AuthorName,
			Email: constant.AuthorEmail,
		}},
		Metadata: make(map[string]any),
	}
}
