package pangu

var _ option = (*optionValidate)(nil)

type optionValidate struct {
	isValidate bool
}

// Validate 配置是否处理默认值
func Validate(isValidate bool) *optionValidate {
	return &optionValidate{
		isValidate: isValidate,
	}
}

func (b *optionValidate) apply(options *options) {
	options.isValidate = b.isValidate
}
