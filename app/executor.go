package app

// Executor Serve命令执行前执行
type Executor interface {
	// Run 运行命令
	Run() (err error)

	// Name 名称
	Name() string

	// Type 类型
	Type() ExecutorType

	// ExecType 执行类型
	ExecType() ExecType
}

// RunExecutors 运行所有执行器
func RunExecutors(executors ...Executor) (err error) {
	for _, executor := range executors {
		if err = executor.Run(); nil != err {
			switch executor.ExecType() {
			case ExecTypeBreak:
				err = nil
				break
			case ExecTypeContinue:
				err = nil
				continue
			case ExecTypeReturn:
				return
			}
		}
	}

	return
}
