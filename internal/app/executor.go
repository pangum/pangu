package app

// Executor Serve命令执行前执行
type Executor interface {
	// Run 运行命令
	Run() (err error)

	// Name 名称
	Name() string

	// Type 类型
	Type() ExecutorType

	// ExecuteType 执行类型
	ExecuteType() ExecuteType
}
