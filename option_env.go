package pangu

var (
	_        = Env
	_ option = (*optionEnv)(nil)
)

type optionEnv struct {
	key   string
	value string
}

// Env 环境变量
func Env(key string, value string) *optionEnv {
	return &optionEnv{
		key:   key,
		value: value,
	}
}

func (e *optionEnv) apply(options *options) {
	options.envs = append(options.envs, &env{
		key:   e.key,
		value: e.value,
	})
}
