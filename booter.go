package boot

import (
	"github.com/harluo/boot/internal/core"
)

// Booter 启动器
// 实现`Commands`方法用于指定应用程序内可以被执行的命令`application.Command`
// 实现`Arguments`方法用于指定应用程序的参数`application.Argument`
type Booter = core.Booter
