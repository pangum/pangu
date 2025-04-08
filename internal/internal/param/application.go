package param

import (
	"github.com/pangum/pangu/internal/internal/constant"
	"github.com/pangum/pangu/internal/internal/core"
	"github.com/pangum/pangu/internal/internal/loader"
)

type Application struct {
	*Config
	*Help
	*Banner
	*Code
	*Timeout

	// 合法性验证，包括
	// 启动器构造方法合法性验证
	// 构造方法合法性验证
	Verify bool
	// 应用描述
	Description string
	// 使用方式
	Usage string

	// 版权
	Copyright string
	// 作者
	Authors core.Authors
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

		Verify:      true,
		Description: "一个使用github.com/pangum/pangu构建的应用程序，可以使用应用程序提供的命令来使用本程序",
		Usage:       "一个使用github.com/pangum/pangu构建的应用程序",

		Copyright: constant.Copyright,
		Authors: core.Authors{{
			Name:  constant.AuthorName,
			Email: constant.AuthorEmail,
		}},
		Metadata: make(map[string]any),
	}
}
