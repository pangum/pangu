package app

const (
	// ExecTypeBreak 中断执行
	ExecTypeBreak ExecType = 1
	// ExecTypeContinue 继续执行
	ExecTypeContinue ExecType = 2
	// ExecTypeReturn 返回
	ExecTypeReturn ExecType = 3
)

type ExecType int
