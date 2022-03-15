package pangu

var (
	_        = DisableDefault
	_ option = (*optionDefault)(nil)
)

type optionDefault struct {
	defaults bool
}

// DisableDefault 配置是否处理默认值
func DisableDefault() *optionDefault {
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
