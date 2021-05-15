package pangu

var _ option = (*optionCopyright)(nil)

type optionCopyright struct {
	copyright string
}

// Copyright 配置版权
func Copyright(copyright string) *optionCopyright {
	return &optionCopyright{
		copyright: copyright,
	}
}

func (b *optionCopyright) apply(options *options) {
	options.copyright = b.copyright
}
