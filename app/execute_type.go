package app

const (
	// ExecuteTypeBreak 中断执行
	ExecuteTypeBreak ExecuteType = 1
	// ExecuteTypeContinue 继续执行
	ExecuteTypeContinue ExecuteType = 2
	// ExecuteTypeReturn 返回
	ExecuteTypeReturn ExecuteType = 3
)

type ExecuteType int
