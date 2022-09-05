package pangu

var (
	_              = ConfigLoader
	_ option       = (*optionConfigLoader)(nil)
	_ configOption = (*optionConfigLoader)(nil)
)

type optionConfigLoader struct {
	loader configLoader
}

// ConfigLoader 配置监控
func ConfigLoader(loader configLoader) *optionConfigLoader {
	return &optionConfigLoader{
		loader: loader,
	}
}

func (cl *optionConfigLoader) apply(options *options) {
	options.loader = cl.loader
}

func (cl *optionConfigLoader) applyConfig(options *configOptions) {
	options.loader = cl.loader
}
