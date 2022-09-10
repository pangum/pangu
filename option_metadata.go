package pangu

var (
	_        = Metadata
	_ option = (*optionMetadata)(nil)
)

type optionMetadata struct {
	key  string
	value any
}

// Metadata 元数据
func Metadata(key string, value any) *optionMetadata {
	return &optionMetadata{
		key:  key,
		value: value,
	}
}

func (m *optionMetadata) apply(options *options) {
	options.metadata[m.key] = m.value
}
