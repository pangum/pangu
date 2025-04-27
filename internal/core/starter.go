package core

// Starter 启动器，全局只能有一个启动器，且只能返回 Starter 才能被正确的启动
type Starter interface {
	// application.Lifecycle

	// Startup 启动
	Startup(application *Application) error
}
