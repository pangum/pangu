package app

// Bootstrap 启动器，全局只能有一个启动器，且只能返回 Bootstrap 才能被正确的启动
type Bootstrap interface {
	// Before 生命周期方法
	// 启动前执行
	Before() error

	// Startup 启动
	Startup() error

	// After 生命周期方法
	// 退出前执行
	After() error
}
