package arg

var (
	_        = Envs
	_        = Environment
	_ option = (*optionEnvs)(nil)
)

type optionEnvs struct {
	envs []string
}

// Envs 环境变量列表
func Envs(envs ...string) *optionEnvs {
	return &optionEnvs{
		envs: envs,
	}
}

// Environment 环境变量列表
func Environment(environment ...string) *optionEnvs {
	return Envs(environment...)
}

func (e *optionEnvs) apply(options *options) {
	options.envs = e.envs
}
