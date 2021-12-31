package pangu

type env struct {
	key   string
	value string
}

// NewEnv 创建环境变量
func NewEnv(key string, value string) *env {
	return &env{
		key:   key,
		value: value,
	}
}
