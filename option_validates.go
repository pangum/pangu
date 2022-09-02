package pangu

var (
	_        = DisableValidate
	_        = DisableValidates
	_ option = (*optionValidate)(nil)
)

type optionValidate struct {
	validates bool
}

// DisableValidate 关闭数据验证
func DisableValidate() *optionValidate {
	return DisableValidates()
}

// DisableValidates 关闭数据验证
func DisableValidates() *optionValidate {
	return &optionValidate{
		validates: false,
	}
}

func (v *optionValidate) apply(options *options) {
	options.validates = v.validates
}
