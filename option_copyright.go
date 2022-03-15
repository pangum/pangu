package pangu

var (
	_        = Copyright
	_ option = (*optionCopyright)(nil)
)

type optionCopyright struct {
	copyright string
}

// Copyright 配置版权
func Copyright(copyright string) *optionCopyright {
	return &optionCopyright{
		copyright: copyright,
	}
}

func (c *optionCopyright) apply(options *options) {
	options.copyright = c.copyright
}
