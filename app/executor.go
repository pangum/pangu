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

// RunExecutors 运行所有执行器
func RunExecutors(executors ...Executor) (err error) {
	for _, executor := range executors {
		if err = executor.Run(); nil != err {
			switch executor.ExecuteType() {
			case ExecuteTypeBreak:
				err = nil
				break
			case ExecuteTypeContinue:
				err = nil
				continue
			case ExecuteTypeReturn:
				return
			}
		}
	}

	return
}
