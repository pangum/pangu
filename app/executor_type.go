package app

const (
	// ExecutorTypeBeforeServe Serve命令执行前执行
	ExecutorTypeBeforeServe ExecutorType = 1
	// ExecutorTypeBeforeAll 所有命令执行前执行
	ExecutorTypeBeforeAll ExecutorType = 2
	// ExecutorTypeAfterServe 在Serve命令执行后执行
	ExecutorTypeAfterServe ExecutorType = 3
	// ExecutorTypeAfterAll 所有命令执行后执行
	ExecutorTypeAfterAll ExecutorType = 4
)

type ExecutorType int
