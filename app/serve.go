package app

// Serve 描述一个服务器，可以是Http服务器，也可以是gRPC服务器或者一个MQ的消费者
// 泛指一个可以长期执行的服务
type Serve interface {
	// Run 运行服务器
	Run() (err error)

	// Name 服务器名称
	Name() string
}
