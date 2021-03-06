package pangu

var (
	_        = Verify
	_        = DisableVerify
	_ option = (*optionVerify)(nil)
)

type optionVerify struct {
	verify bool
}

// Verify 启用合法性验证
func Verify() *optionVerify {
	return &optionVerify{
		verify: true,
	}
}

// DisableVerify 禁用合法性验证
func DisableVerify() *optionVerify {
	return &optionVerify{
		verify: false,
	}
}

func (v *optionVerify) apply(options *options) {
	options.verify = v.verify
}
