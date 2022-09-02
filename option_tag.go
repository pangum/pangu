package pangu

var (
	_ = Tag
	_ = DefaultTag

	_ option = (*optionTag)(nil)
)

type optionTag struct {
	defaults string
}

// Tag 标签
func Tag(defaults string) *optionTag {
	return &optionTag{
		defaults: defaults,
	}
}

// DefaultTag 默认标签值
func DefaultTag(tag string) *optionTag {
	return &optionTag{
		defaults: tag,
	}
}

func (t *optionTag) apply(options *options) {
	if `` != t.defaults {
		options.tag.defaults = t.defaults
	}
}
