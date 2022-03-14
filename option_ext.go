package pangu

var (
	_        = Ext
	_ option = (*optionExt)(nil)
)

type optionExt struct {
	ext     string
	options *configOptions
}

// Ext 配置扩展名
func Ext(ext string, opts ...configOption) *optionExt {
	_options := defaultConfigOptions()
	for _, opt := range opts {
		opt.applyConfig(_options)
	}

	return &optionExt{
		ext:     ext,
		options: _options,
	}
}

func (cp *optionExt) apply(options *options) {
	switch cp.options.typ {
	case ConfigTypeAppend:
		options.extensions = append(options.extensions, cp.ext)
	case ConfigTypeOverride:
		options.extensions = []string{cp.ext}
	}
}
