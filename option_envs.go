package pangu

var _ option = (*optionEnvs)(nil)

type optionEnvs struct {
	envs []*env
}

// Envs 配置环境变量
func Envs(envs ...*env) *optionEnvs {
	return &optionEnvs{
		envs: envs,
	}
}

func (c *optionEnvs) apply(options *options) {
	options.envs = c.envs
}
