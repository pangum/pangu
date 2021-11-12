package pangu

var _ option = (*optionDefault)(nil)

type optionDefault struct {
	_default bool
}

// DisableDefault 配置是否处理默认值
func DisableDefault() *optionDefault {
	return &optionDefault{
		_default: false,
	}
}

func (d *optionDefault) apply(options *options) {
	options._default = d._default
}
