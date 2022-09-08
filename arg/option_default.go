package arg

var (
	_        = Default
	_ option = (*optionDefault)(nil)
)

type optionDefault struct {
	_default any
}

// Default 默认值
func Default(_default any) *optionDefault {
	return &optionDefault{
		_default: _default,
	}
}

func (d *optionDefault) apply(options *options) {
	options._default = d._default
}
