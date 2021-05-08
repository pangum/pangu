package pangu

var _ option = (*optionDefault)(nil)

type optionDefault struct {
	isDefault bool
}

// Default 配置是否处理默认值
func Default(isDefault bool) *optionDefault {
	return &optionDefault{
		isDefault: isDefault,
	}
}

func (b *optionDefault) apply(options *options) {
	options.isDefault = b.isDefault
}
