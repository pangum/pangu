package arg

var (
	_        = String
	_        = Bool
	_        = Int
	_        = Int64
	_        = Float64
	_ option = (*optionDefault)(nil)
)

type optionDefault struct {
	_default any
}

// String 默认值
func String(value *string) *optionDefault {
	return _default(value)
}

// Bool 默认值
func Bool(value *bool) *optionDefault {
	return _default(value)
}

// Int 默认值
func Int(value *int) *optionDefault {
	return _default(value)
}

// Int64 默认值
func Int64(value *int64) *optionDefault {
	return _default(value)
}

// Float64 默认值
func Float64(value *float64) *optionDefault {
	return _default(value)
}

func _default(_default any) *optionDefault {
	return &optionDefault{
		_default: _default,
	}
}

func (d *optionDefault) apply(options *options) {
	options._default = d._default
}
