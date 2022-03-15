package pangu

var (
	_        = DisableDefault
	_        = DisableDefaults
	_ option = (*optionDefault)(nil)
)

type optionDefault struct {
	defaults bool
}

// DisableDefault 配置是否处理默认值
func DisableDefault() *optionDefault {
	return DisableDefaults()
}

// DisableDefaults 配置是否处理默认值
func DisableDefaults() *optionDefault {
	return &optionDefault{
		defaults: false,
	}
}

func (d *optionDefault) apply(options *options) {
	options.defaults = d.defaults
}

func (d *optionDefault) applyConfig(options *configOptions) {
	options.defaults = d.defaults
}
