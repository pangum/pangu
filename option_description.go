package pangu

var _ option = (*optionDescription)(nil)

type optionDescription struct {
	description string
}

// Description 配置应用描述
func Description(description string) *optionDescription {
	return &optionDescription{
		description: description,
	}
}

func (b *optionDescription) apply(options *options) {
	options.description = b.description
}
