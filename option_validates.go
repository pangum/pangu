package pangu

var (
	_        = DisableValidate
	_        = DisableValidates
	_ option = (*optionValidate)(nil)
)

type optionValidate struct {
	validates bool
}

// DisableValidate 配置是否处理默认值
func DisableValidate() *optionValidate {
	return DisableValidates()
}

// DisableValidates 配置是否处理默认值
func DisableValidates() *optionValidate {
	return &optionValidate{
		validates: false,
	}
}

func (v *optionValidate) apply(options *options) {
	options.validates = v.validates
}
