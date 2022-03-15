package pangu

var (
	_              = Ext
	_ option       = (*optionExt)(nil)
	_ configOption = (*optionExt)(nil)
)

type optionExt struct {
	ext     string
	options *optionOptions
}

// Ext 配置扩展名
func Ext(ext string, opts ...optionOption) *optionExt {
	_options := defaultOptionOptions()
	for _, opt := range opts {
		opt.applyOption(_options)
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

func (cp *optionExt) applyConfig(options *configOptions) {
	switch cp.options.typ {
	case ConfigTypeAppend:
		options.extensions = append(options.extensions, cp.ext)
	case ConfigTypeOverride:
		options.extensions = []string{cp.ext}
	}
}
