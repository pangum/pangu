package pangu

var _ option = (*optionName)(nil)

type optionName struct {
	name string
}

// Name 配置应用名称
func Name(name string) *optionName {
	return &optionName{
		name: name,
	}
}

func (b *optionName) apply(options *options) {
	options.name = b.name
}
