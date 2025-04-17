package core

import (
	"github.com/harluo/boot/internal/application"
)

// Startup 启动器，全局只能有一个启动器，且只能返回 Startup 才能被正确的启动
type Startup interface {
	application.Lifecycle

	// Startup 启动
	Startup(application *Application) error
}
