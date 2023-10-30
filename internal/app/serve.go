package app

// Serve 描述一个服务器，可以是Http服务器，也可以是gRPC服务器或者一个MQ的消费者
// 泛指一个可以长期执行的服务
type Serve interface {
	Stopper

	// Before 生命周期方法
	// 启动前执行
	Before() error

	// Start 运行服务
	Start() (err error)

	// Name 服务名称
	Name() string

	// After 生命周期方法
	// 退出前执行
	After() error
}
