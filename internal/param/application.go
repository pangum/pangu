package param

import (
	"github.com/pangum/pangu/internal/constant"
	"github.com/pangum/pangu/internal/internal/app"
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
	Authors app.Authors
	// 元数据
	Metadata map[string]any
	Set      bool
}

func NewApplication() *Application {
	return &Application{
		Config:  newConfig(),
		Help:    newHelp(),
		Banner:  newBanner(),
		Code:    newCode(),
		Timeout: NewTimeout(),

		Verify:      true,
		Description: "一个使用github.com/pangum/pangu构建的应用程序，可以使用应用程序提供的命令来使用本程序",
		Usage:       "一个使用github.com/pangum/pangu构建的应用程序",

		Copyright: constant.Copyright,
		Authors: app.Authors{{
			Name:  constant.AuthorName,
			Email: constant.AuthorEmail,
		}},
		Metadata: make(map[string]any),
	}
}

func (a *Application) Override(latest *Application) {
	if latest.Config.Set {
		a.Config = latest.Config
	}
	if latest.Help.Set {
		a.Help = latest.Help
	}
	if latest.Banner.Set {
		a.Banner = latest.Banner
	}
	if latest.Code.Set {
		a.Code = latest.Code
	}
	if latest.Set {
		a.Verify = latest.Verify
		a.Description = latest.Description
		a.Usage = latest.Usage
		a.Copyright = latest.Copyright
		a.Authors = latest.Authors
		a.Metadata = latest.Metadata
	}
}
