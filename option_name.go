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

func (n *optionName) apply(options *options) {
	options.name = n.name
}

func (n *optionName) applyProvide(options *provideOptions) {
	options.name = n.name
}

func (n *optionName) applyInvoke(options *invokeOptions) {
	options.name = n.name
}
