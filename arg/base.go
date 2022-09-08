package arg

type base struct {
	// 名称
	name string
	// 默认值
	_default any
	// 绑定
	destination any
	// 别名
	aliases []string
	// 使用方法
	usage string
	// 环境变量列表
	envs []string
	// 是否是必须
	required bool
	// 是否隐藏
	hidden bool
	// 默认显示字符串
	dt string
}

func _new(name string, opts ...option) *base {
	_options := defaultOption()
	for _, opt := range opts {
		opt.apply(_options)
	}

	return &base{
		name:        name,
		_default:    _options._default,
		destination: _options.destination,
		aliases:     _options.aliases,
		usage:       _options.usage,
		envs:        _options.envs,
		required:    _options.required,
		hidden:      _options.hidden,
		dt:          _options.dt,
	}
}

func (b *base) Name() string {
	return b.name
}

func (b *base) Default() any {
	return b._default
}

func (b *base) Destination() any {
	return b.destination
}

func (b *base) Aliases() []string {
	return b.aliases
}

func (b *base) Usage() string {
	return b.usage
}

func (b *base) DefaultText() string {
	return b.dt
}

func (b *base) Required() bool {
	return b.required
}

func (b *base) Hidden() bool {
	return b.hidden
}
