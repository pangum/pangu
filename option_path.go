package pangu

var (
	_              = Path
	_ option       = (*optionPath)(nil)
	_ configOption = (*optionPath)(nil)
)

type optionPath struct {
	path    string
	options *optionOptions
}

// Path 配置路径
func Path(path string, opts ...optionOption) *optionPath {
	_options := defaultOptionOptions()
	for _, opt := range opts {
		opt.applyOption(_options)
	}

	return &optionPath{
		path:    path,
		options: _options,
	}
}

func (cp *optionPath) apply(options *options) {
	switch cp.options.typ {
	case ConfigTypeAppend:
		options.paths = append(options.paths, cp.path)
	case ConfigTypeOverride:
		options.paths = []string{cp.path}
	}
}

func (cp *optionPath) applyConfig(options *configOptions) {
	switch cp.options.typ {
	case ConfigTypeAppend:
		options.paths = append(options.paths, cp.path)
	case ConfigTypeOverride:
		options.paths = []string{cp.path}
	}
}
