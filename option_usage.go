package pangu

var _ option = (*optionUsage)(nil)

type optionUsage struct {
	usage string
}

// Usage 配置应用描述
func Usage(usage string) *optionUsage {
	return &optionUsage{
		usage: usage,
	}
}

func (b *optionUsage) apply(options *options) {
	options.usage = b.usage
}
