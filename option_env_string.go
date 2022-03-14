package pangu

var (
	_        = StringEnv
	_ option = (*optionStringEnv)(nil)
)

type optionStringEnv struct {
	env string
}

// StringEnv 环境变量列表
func StringEnv(env string) *optionStringEnv {
	return &optionStringEnv{
		env: env,
	}
}

func (se *optionStringEnv) apply(options *options) {
	options.envs = append(options.envs, parseEnv(se.env))
}
