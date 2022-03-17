package pangu

var (
	_        = Env
	_        = Environment
	_ option = (*optionEnvironment)(nil)
)

type optionEnvironment struct {
	key   string
	value string
}

// Env 环境变量
func Env(key string, value string) *optionEnvironment {
	return Environment(key, value)
}

// Environment 环境变量
func Environment(key string, value string) *optionEnvironment {
	return &optionEnvironment{
		key:   key,
		value: value,
	}
}

func (e *optionEnvironment) apply(options *options) {
	options.environments = append(options.environments, &environment{
		key:   e.key,
		value: e.value,
	})
}
