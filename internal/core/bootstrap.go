package core

// Bootstrap 启动器，全局只能有一个启动器，且只能返回 Bootstrap 才能被正确的启动，需要完成的事情有
// 添加系统的依赖关系
// 使用 pangu.Application.Provide 添加 pangu.Serve 实现类到 pangu.Application 中
// 使用 pangu.Application.Provide 添加 pangu.Base 实现类到 pangu.Application 中
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
