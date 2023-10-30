package core

import (
	"github.com/pangum/pangu/internal/app"
)

// Bootstrap 启动器，全局只能有一个启动器，且只能返回 Bootstrap 才能被正确的启动
type Bootstrap interface {
	app.Lifecycle

	// Startup 启动
	Startup(app *Application) error
}
