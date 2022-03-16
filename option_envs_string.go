package pangu

var (
	_        = StringEnv
	_        = StringEnvs
	_ option = (*optionStringEnvs)(nil)
)

type optionStringEnvs struct {
	envs []string
}

// StringEnv 环境变量
func StringEnv(env string) *optionStringEnvs {
	return &optionStringEnvs{
		envs: []string{env},
	}
}

// StringEnvs 环境变量列表
func StringEnvs(envs ...string) *optionStringEnvs {
	return &optionStringEnvs{
		envs: envs,
	}
}

func (se *optionStringEnvs) apply(options *options) {
	for _, _env := range se.envs {
		options.envs = append(options.envs, parseEnv(_env))
	}
}
