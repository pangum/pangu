package pangu

var (
	_ = Named

	_ option        = (*optionNamed)(nil)
	_ provideOption = (*optionNamed)(nil)
	_ invokeOption  = (*optionNamed)(nil)
)

type optionNamed struct {
	name string
}

// Named 应用名称
func Named(name string) *optionNamed {
	return &optionNamed{
		name: name,
	}
}

func (n *optionNamed) apply(_ *options) {
	App = n.name
}

func (n *optionNamed) applyProvide(options *provideOptions) {
	options.name = n.name
}

func (n *optionNamed) applyInvoke(options *invokeOptions) {
	options.name = n.name
}
