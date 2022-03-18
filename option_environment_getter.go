package pangu

var (
	_              = EnvironmentGetter
	_ option       = (*optionEnvironmentGetter)(nil)
	_ configOption = (*optionEnvironmentGetter)(nil)
)

type optionEnvironmentGetter struct {
	getter environmentGetter
}

// EnvironmentGetter 环境变量
func EnvironmentGetter(getter environmentGetter) *optionEnvironmentGetter {
	return &optionEnvironmentGetter{
		getter: getter,
	}
}

func (eg *optionEnvironmentGetter) apply(options *options) {
	options.environmentGetter = eg.getter
}

func (eg *optionEnvironmentGetter) applyConfig(options *configOptions) {
	options.environmentGetter = eg.getter
}
