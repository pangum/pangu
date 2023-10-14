package pangu

import (
	"github.com/pangum/pangu/internal/core"
)

// Application 应用程序，可以加入两种种类型的程序
// Serve 用于描述应用程序内的服务
// Command 用于描述应用程序内可以被执行的命令
type Application = core.Application
