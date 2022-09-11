package cmd

var (
	_        = Usage
	_ option = (*optionUsage)(nil)
)

type optionUsage struct {
	usage string
}

// Usage 命令使用描述
func Usage(usage string) *optionUsage {
	return &optionUsage{
		usage: usage,
	}
}

func (u *optionUsage) apply(options *options) {
	options.usage = u.usage
}
