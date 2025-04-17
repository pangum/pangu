package boot

import (
	"github.com/harluo/boot/internal/builder"
	"github.com/harluo/boot/internal/core"
)

// Application 应用程序，可以加入两种种类型的程序
// application.Serve 用于描述应用程序内的服务
// application.Command 用于描述应用程序内可以被执行的命令
// application.Argument 用于描述应用程序的参数
type Application = core.Application

func New() *builder.Application {
	return builder.NewApplication()
}
