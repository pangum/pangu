package pangu

import (
	"github.com/pangum/pangu/internal/builder"
	"github.com/pangum/pangu/internal/core"
)

var _ = New

// Application 应用程序，可以加入两种种类型的程序
// app.Serve 用于描述应用程序内的服务
// app.Command 用于描述应用程序内可以被执行的命令
// app.Argument 用于描述应用程序的参数
type Application = core.Application

func New() *builder.Application {
	return builder.NewApplication()
}
