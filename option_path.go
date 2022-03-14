package pangu

var (
	_        = Path
	_ option = (*optionPath)(nil)
)

type optionPath struct {
	path    string
	options *configOptions
}

// Path 配置路径
func Path(path string, opts ...configOption) *optionPath {
	_options := defaultConfigOptions()
	for _, opt := range opts {
		opt.applyConfig(_options)
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
