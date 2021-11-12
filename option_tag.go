package pangu

var _ option = (*optionTag)(nil)

type optionTag struct {
	_default string
}

// Tag 配置标签
func Tag(_default string) *optionTag {
	return &optionTag{
		_default: _default,
	}
}

// DefaultTag 配置默认值标签
func DefaultTag(tag string) *optionTag {
	return &optionTag{
		_default: tag,
	}
}

func (t *optionTag) apply(options *options) {
	if `` != t._default {
		options.tag._default = t._default
	}
}
