package config

type Watcher interface {
	// OnChanged 配置发生改变时
	OnChanged(path string, from any, to any)

	// OnDeleted 配置被删除时
	OnDeleted(path string)

	// OnError 发生错误时
	OnError(path string, err error)
}
