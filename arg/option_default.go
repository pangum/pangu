package arg

import (
	"time"
)

var (
	_        = String
	_        = StringSlice
	_        = Bool
	_        = BoolSlice
	_        = Int
	_        = IntSlice
	_        = Int64
	_        = Int64Slice
	_        = Float64
	_        = Float64Slice
	_        = Duration
	_        = DurationSlice
	_ option = (*optionDefault)(nil)
)

type optionDefault struct {
	_default any
}

// String 字符串默认值
func String(value string) *optionDefault {
	return _default(value)
}

// StringSlice 字符串列表默认值
func StringSlice(values []string) *optionDefault {
	return _default(values)
}

// Bool 默认值
func Bool(value bool) *optionDefault {
	return _default(value)
}

// BoolSlice 默认值
func BoolSlice(values []bool) *optionDefault {
	return _default(values)
}

// Int 默认值
func Int(value int) *optionDefault {
	return _default(value)
}

// IntSlice 默认值
func IntSlice(values []int) *optionDefault {
	return _default(values)
}

// Int64 默认值
func Int64(value int64) *optionDefault {
	return _default(value)
}

// Int64Slice 默认值
func Int64Slice(values []int64) *optionDefault {
	return _default(values)
}

// Float64 默认值
func Float64(value float64) *optionDefault {
	return _default(value)
}

// Float64Slice 默认值
func Float64Slice(values []float64) *optionDefault {
	return _default(values)
}

// Duration 默认值
func Duration(value time.Duration) *optionDefault {
	return _default(value)
}

// DurationSlice 默认值
func DurationSlice(values []time.Duration) *optionDefault {
	return _default(values)
}

func _default(_default any) *optionDefault {
	return &optionDefault{
		_default: _default,
	}
}

func (d *optionDefault) apply(options *options) {
	options._default = d._default
}
