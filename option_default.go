package pangu

var _ option = (*optionDefault)(nil)

type optionDefault struct {
	isDefault bool
}

// DisableDefault 配置是否处理默认值
func DisableDefault() *optionDefault {
	return &optionDefault{
		isDefault: false,
	}
}

func (b *optionDefault) apply(options *options) {
	options.isDefault = b.isDefault
}
