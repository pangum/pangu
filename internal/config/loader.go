package config

type Loader[T any] interface {
	// Load 加载配置
	Load(path string) (value *T, err error)
}
