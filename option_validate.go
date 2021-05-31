package pangu

var _ option = (*optionValidate)(nil)

type optionValidate struct {
	isValidate bool
}

// DisableValidate 配置是否处理默认值
func DisableValidate() *optionValidate {
	return &optionValidate{
		isValidate: false,
	}
}

func (b *optionValidate) apply(options *options) {
	options.isValidate = b.isValidate
}
