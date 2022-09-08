package arg

var (
	_        = Usage
	_ option = (*optionUsage)(nil)
)

type optionUsage struct {
	usage string
}

// Usage 参数描述
func Usage(usage string) *optionUsage {
	return &optionUsage{
		usage: usage,
	}
}

func (u *optionUsage) apply(options *options) {
	options.usage = u.usage
}
