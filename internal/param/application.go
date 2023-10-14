package param

import (
	"github.com/pangum/pangu/internal/internal"
)

type Application struct {
	*Config
	*Help
	*Banner

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
	Authors internal.Authors
	// 元数据
	Metadata map[string]any
}

func NewApplication() *Application {
	return &Application{
		Config: NewConfig(),
		Help:   NewHelp(),
	}
}

func (a *Application) Override(application *Application) {}
