package pangu

var _ option = (*optionValidate)(nil)

type optionValidate struct {
	validate bool
}

// DisableValidate 配置是否处理默认值
func DisableValidate() *optionValidate {
	return &optionValidate{
		validate: false,
	}
}

func (b *optionValidate) apply(options *options) {
	options.validate = b.validate
}
