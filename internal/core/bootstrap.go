package core

import (
	"github.com/pangum/core/internal/application"
)

// Bootstrap 启动器，全局只能有一个启动器，且只能返回 Bootstrap 才能被正确的启动
type Bootstrap interface {
	application.Lifecycle

	// Startup 启动
	Startup(app *Application) error
}
